# Digo

###Dependency Injeccion for Go

Digo is inspired in the way that the Spring framework for Java allows dependency injection using an configuration file. The difference is that digo uses Json instead of XML and, for language limitations, the Types must be registered.

##Install

Download the repository with:
``go get github.com/cone/digo``

Then include the project inside your project like this:
``import "github.com/cone/digo"``

##What it does

Digo allows dependency injection using a configuration file.

Lets suppose we have ProductController struct as follows:

    type ProductController struct{
        DB Repository
        Validator DataValidator
    }
    
``ProductController`` has 2 dependencies: ``DB`` of type ``Repository`` and ``Validator`` of type ``DataValidator``, we want to follow the inversion of controll principle so both are interfaces:

    type Repository interface{
      Get(int) interface{}
    }
    type DataValidator interface{
      Validate(interface{}) bool
    }
    
Also we have custom implementations of each of those interfaces for products, for example, we have these for the ``Repository`` interface:

    type ProductRepo struct{}
    func (this *ProductRepo) Get(id int) interface{}{
      //...reads from db...
      return product
    }
    
    type TestProductRepo struct{}
    func (this TestProductRepo) Get(id int) interface{}{
      return Product{id, "Test Product"}
    }
    
First we have to add the types to the ``TypeRegister`` like this:

    TypeRegistry.Add(ProductController{})
    TypeRegistry.Add(ProductRepo{})
    TypeRegistry.Add(ProductValidator{})
    
Then we should write a configuration file like this called ``prod.json``:

    {
      nodes: {
        "product_repo": {
          "type": "myPackage.ProductRepo",
          "is_pointer": true,
        },
        "product_validator": {
          "type": "myPackage.ProductValidator",
        }
        "product_ctrl": {
          "type": "myPackage.ProductController",
          "deps": [
            {
              "id": "product_repo",
              "field": "DB"
            },
            {
              "id": "product_validator",
              "field": "Validator"
            }
          ]
        }
      }
    }
    
In the file we define the types and their dependencies. We use an alias to name the type e.g. ``product_repo`` for ``myPackage.ProductRepo`` then we specify the ``type`` and if the dependency is a pointer we set the ``is_pointer`` field to ``true``, note that this is true for ``myPackage.ProductRepo`` because only a pointer to it satisfies the ``Repository`` interface.
    
Now we can create a ``Context`` with that file. The ``Context`` allow us to easily change the dependencies whe are injecting, that is useful if when testing: suppose we have a ``test.json`` file, in that file, instead of ``myPackage.ProductRepo`` we specify ``myPackage.TestProductRepo`` which is a test double. having those two files we can create a production context and a test context.

    prodCtx, err := Digo.Context("prod.json")
    //...error handling...
    
    testCtx, err := Digo.Context("test.json")
    //...error handling...

Finally the context is able to generate our ``ProductController`` using the alias we specified in the json file:

    pi, err := prodCtx.Get("product_repo")
    //...error handling...
    productController := pi.(ProductController)
    
    ti, err := testCtx.Get("product_repo")
    //...error handling...
    testProductController := ti.(ProductController)
    
The gererated ``ProductController`` will already have its dependencies ``ProductRepo`` and ``ProductValidator`` ready for use!

    product := productController.DB.Get(1)

##Usage

###TypeRegistry

Its pourpose is to register the types of the elements that are going to be injected or created.

####Add

It allows to add a type. You must to pass a pointer to it, that is must be specified in the config file with the ``is_pointer`` flag.

    TypeRegistry.Add(ProductController{})
    
####AddType

It allows to add a type passing directly a reflect.Type element.

    TypeRegistry.AddType(reflect.TypeOf(PorductController{}))
    
####Get

It allows to get a reflect.Type element passing a string containing the name of the type.

    t, _ := TypeRegistry.Get("myPackage")
    
###Digo

It Manages the different contexts we may have.

####Context

It returns a Context element from a given config file.

    ctx, _ := Digo.Context("prod.json")
    
###Context

It holds a given configuration that will be use to determine what dependencies will be injected to an element

####Get

It will return a copy on the given type by its alias. The dependencies of pointer type will be shared.

    type Foo interface{
        GetDBConnection() DB    
    }
    
    type Bar struct{
        db DB
    }
    func (this *Bar) GetDBConnection() DB{
        return this.db
    }
    
    type Baz struct{
        Msg string
        Con Foo
    }
    
    /*in config file
    ...
    "bar":{
        "type": "myPackage.Bar",
        "is_pointer": true
    },
    "baz":{
        "type": "myPackage.Baz",
        "deps": [
            {
                "id": "bar",
                "field": "Con"
            }
        ]
    }
    ...
    */
    
    i1, _ := ctx.Get("baz")
    i2, _ := ctx.Get("baz")
    
    baz1 := i1.(Baz)
    baz2 := i2.(Baz)
    
    baz1.Msg = "Hello"
    baz2.Msg = "World"
    
    //baz1.Msg != baz2.Msg but baz1.Con is the same as baz2.Con, it is "shared"
    
####Copy

It returns a deep copy of the specified type by its alias. The copies are totally independient from each other.

    i1, _ := ctx.Copy("baz")
    i2, _ := ctx.Copy("baz")
    
    bar1 := i1.(Baz)
    bar2 := i2.(Baz)
    
    //baz1.Msg != baz2.Msg and baz1.Con != baz2.Con

####Single

It returns a pointer of a given type by its alias, so it behaves as a ``singleton``. Type assertion should be made against a pointer to the corresponding type.

    i1, _ := ctx.Single("baz")
    i2, _ := ctx.Single("baz")
    
    bar1 := i1.(*Baz)
    bar2 := i2.(*Baz)
    
    //bar1 and bar2 are pointers to the same object
    
##Initializer Interface

If one of the dependencies implements the ``Initializer`` interface, th ``BeforeInject`` function will be called before the dependency is injected. This could be useful to initialize the values of the dependency e.g.

    type Bar struct{
        db DB
    }
    func (this *Bar) GetDBConnection() DB{
        return this.db
    }
    func (this *Bar) BeforeInject() error{
        return this.db.Open()
    }
    
## Contributing

1. Create your feature branch (`git checkout -b feature/my-new-feature`)
2. Commit your changes (`git commit -am 'Add some feature'`)
3. Push to the branch (`git push origin feature/my-new-feature`)
4. Create a new Pull Request
