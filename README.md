# digo

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
    func Get(id int) interface{}{
      //...reads from db...
      return product
    }
    
    type TestProductRepo struct{}
    func Get(id int) interface{}{
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
    
Now we can create a ``Context`` with that file. The ``Context`` allow us to easily change the dependencies whe are injecting, that is useful if when testing: suppose we have a ``test.json`` file, in that file, instead of ``myPackage.ProductRepo`` we specify ``myPackage.TestProductRepo`` which is a test double. having those two files we can create a production context and a test context.

    prodCtx, err := Digo.Context("prod.json")
    //...error handling...
    
    testCtx, err := Digo.Context("test.json")
    //...error handling...

Finally the context is able to generate our ``ProductController`` using the alias we specified in the json file:

    productController, err := prodCtx.Get("product_repo")
    //...error handling...
    
    testProductController, err := testCtx.Get("product_repo")
    //...error handling...
    
The gererated ``ProductController`` will already have its dependencies ``ProductRepo`` and ``ProductValidator`` ready for use!

    product := productController.DB.Get(1)
