# Digo

[![Build Status](https://travis-ci.org/cone/digo.svg?branch=master)](https://travis-ci.org/cone/digo)
[![views](https://sourcegraph.com/api/repos/github.com/cone/digo/.counters/views.svg)](https://sourcegraph.com/github.com/cone/digo)

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

    prodCtx, err := ContextFor("prod.json")
    //...error handling...
    
    testCtx, err := ContextFor("test.json")
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

###ContextFor

It returns a Context element from a given config file.

    ctx, _ := ContextFor("prod.json")

###TypeRegistry

Its pourpose is to register the types of the elements that are going to be injected or created.

####Add

It allows to add a type. You must not pass a pointer to it, that is must be specified in the config file with the ``is_pointer`` flag.

    TypeRegistry.Add(ProductController{})
    
####AddType

It allows to add a type passing directly a reflect.Type element.

    TypeRegistry.AddType(reflect.TypeOf(PorductController{}))
    
####Get

It allows to get a reflect.Type element passing a string containing the name of the type.

    t, _ := TypeRegistry.Get("myPackage.Foo")
    
###Context

It holds a given configuration that will be use to determine what dependencies will be injected to an element

####Get

It will return the given type's interface by its alias.

    i, _ := cxt.Get("foo")
    foo := i.(Foo)
    
##Configuration file

The configuration file holds a Json string describing the components and their dependencies.

###Component fields

####Type

It describes the type of the component including the package.

    ...
    "type": "myPackage.Foo",
    ...
    
####Is_pointer

A flag to specify that the component is a pointer (false by default). If it is not a "singleton", digo will create a new instance and return a pointer to it.

    ....
    "is_pointer": true,
    ....
    
####Scope

It descibes the scope of the component. This field can hold two values: prototype and singleton.

If 'prototype' is used, a copy of the component will be created each time Context.Get is called. This is the default behavior is no 'scope' is specified.

    ...
    "scope": "prototype"
    ...
    
If 'singleton' is used, a reference to the same component will be returned each time Context.Get is called. The variable or struct field that will hold the singleton should be a pointer.

    ...
    "scope": "singleton"
    ...

####Deps

It descibes the dependencies of the component. Here the id of the component used as a dependency should be specified and the name of the field that will hold it e.g.

    ...
    "foo":{
        "type": "myPackage.Foo"
    },
    "deps":[
        {
          "id": "foo",
          "field": "MyFoo"
        }
      ]
    ...

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
