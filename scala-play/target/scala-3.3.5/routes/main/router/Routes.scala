// @GENERATOR:play-routes-compiler
// @SOURCE:conf/routes

package router

import play.core.routing._
import play.core.routing.HandlerInvokerFactory._

import play.api.mvc._

import _root_.controllers.Assets.Asset

class Routes(
  override val errorHandler: play.api.http.HttpErrorHandler, 
  // @LINE:1
  ProductsController_0: javax.inject.Provider[controllers.ProductsController],
  // @LINE:7
  CategoriesController_1: javax.inject.Provider[controllers.CategoriesController],
  // @LINE:13
  CartController_2: javax.inject.Provider[controllers.CartController],
  val prefix: String
) extends GeneratedRouter {

  @javax.inject.Inject()
  def this(errorHandler: play.api.http.HttpErrorHandler,
    // @LINE:1
    ProductsController_0: javax.inject.Provider[controllers.ProductsController],
    // @LINE:7
    CategoriesController_1: javax.inject.Provider[controllers.CategoriesController],
    // @LINE:13
    CartController_2: javax.inject.Provider[controllers.CartController]
  ) = this(errorHandler, ProductsController_0, CategoriesController_1, CartController_2, "/")

  def withPrefix(addPrefix: String): Routes = {
    val prefix = play.api.routing.Router.concatPrefix(addPrefix, this.prefix)
    router.RoutesPrefix.setPrefix(prefix)
    new Routes(errorHandler, ProductsController_0, CategoriesController_1, CartController_2, prefix)
  }

  private val defaultPrefix: String = {
    if (this.prefix.endsWith("/")) "" else "/"
  }

  def documentation = List(
    ("""GET""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """products""", """@controllers.ProductsController@.listProducts"""),
    ("""GET""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """products/""" + "$" + """id<[^/]+>""", """@controllers.ProductsController@.getProduct(id:Int)"""),
    ("""POST""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """products""", """@controllers.ProductsController@.addProduct"""),
    ("""PUT""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """products/""" + "$" + """id<[^/]+>""", """@controllers.ProductsController@.updateProduct(id:Int)"""),
    ("""DELETE""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """products/""" + "$" + """id<[^/]+>""", """@controllers.ProductsController@.deleteProduct(id:Int)"""),
    ("""GET""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """categories""", """@controllers.CategoriesController@.listCategories"""),
    ("""GET""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """categories/""" + "$" + """id<[^/]+>""", """@controllers.CategoriesController@.getCategory(id:Int)"""),
    ("""POST""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """categories""", """@controllers.CategoriesController@.addCategory"""),
    ("""PUT""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """categories/""" + "$" + """id<[^/]+>""", """@controllers.CategoriesController@.updateCategory(id:Int)"""),
    ("""DELETE""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """categories/""" + "$" + """id<[^/]+>""", """@controllers.CategoriesController@.deleteCategory(id:Int)"""),
    ("""GET""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """cart""", """@controllers.CartController@.listCartItems"""),
    ("""POST""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """cart""", """@controllers.CartController@.addCartItem"""),
    ("""PUT""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """cart/""" + "$" + """productId<[^/]+>""", """@controllers.CartController@.updateCartItem(productId:Int)"""),
    ("""DELETE""", this.prefix + (if(this.prefix.endsWith("/")) "" else "/") + """cart/""" + "$" + """productId<[^/]+>""", """@controllers.CartController@.deleteCartItem(productId:Int)"""),
    Nil
  ).foldLeft(Seq.empty[(String, String, String)]) { (s,e) => e.asInstanceOf[Any] match {
    case r @ (_,_,_) => s :+ r.asInstanceOf[(String, String, String)]
    case l => s ++ l.asInstanceOf[List[(String, String, String)]]
  }}


  // @LINE:1
  private lazy val controllers_ProductsController_listProducts0_route = Route("GET",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("products")))
  )
  private lazy val controllers_ProductsController_listProducts0_invoker = createInvoker(
    ProductsController_0.get.listProducts,
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.ProductsController",
      "listProducts",
      Nil,
      "GET",
      this.prefix + """products""",
      """""",
      Seq()
    )
  )

  // @LINE:2
  private lazy val controllers_ProductsController_getProduct1_route = Route("GET",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("products/"), DynamicPart("id", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_ProductsController_getProduct1_invoker = createInvoker(
    ProductsController_0.get.getProduct(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.ProductsController",
      "getProduct",
      Seq(classOf[Int]),
      "GET",
      this.prefix + """products/""" + "$" + """id<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:3
  private lazy val controllers_ProductsController_addProduct2_route = Route("POST",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("products")))
  )
  private lazy val controllers_ProductsController_addProduct2_invoker = createInvoker(
    ProductsController_0.get.addProduct,
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.ProductsController",
      "addProduct",
      Nil,
      "POST",
      this.prefix + """products""",
      """""",
      Seq()
    )
  )

  // @LINE:4
  private lazy val controllers_ProductsController_updateProduct3_route = Route("PUT",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("products/"), DynamicPart("id", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_ProductsController_updateProduct3_invoker = createInvoker(
    ProductsController_0.get.updateProduct(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.ProductsController",
      "updateProduct",
      Seq(classOf[Int]),
      "PUT",
      this.prefix + """products/""" + "$" + """id<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:5
  private lazy val controllers_ProductsController_deleteProduct4_route = Route("DELETE",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("products/"), DynamicPart("id", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_ProductsController_deleteProduct4_invoker = createInvoker(
    ProductsController_0.get.deleteProduct(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.ProductsController",
      "deleteProduct",
      Seq(classOf[Int]),
      "DELETE",
      this.prefix + """products/""" + "$" + """id<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:7
  private lazy val controllers_CategoriesController_listCategories5_route = Route("GET",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("categories")))
  )
  private lazy val controllers_CategoriesController_listCategories5_invoker = createInvoker(
    CategoriesController_1.get.listCategories,
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CategoriesController",
      "listCategories",
      Nil,
      "GET",
      this.prefix + """categories""",
      """""",
      Seq()
    )
  )

  // @LINE:8
  private lazy val controllers_CategoriesController_getCategory6_route = Route("GET",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("categories/"), DynamicPart("id", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_CategoriesController_getCategory6_invoker = createInvoker(
    CategoriesController_1.get.getCategory(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CategoriesController",
      "getCategory",
      Seq(classOf[Int]),
      "GET",
      this.prefix + """categories/""" + "$" + """id<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:9
  private lazy val controllers_CategoriesController_addCategory7_route = Route("POST",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("categories")))
  )
  private lazy val controllers_CategoriesController_addCategory7_invoker = createInvoker(
    CategoriesController_1.get.addCategory,
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CategoriesController",
      "addCategory",
      Nil,
      "POST",
      this.prefix + """categories""",
      """""",
      Seq()
    )
  )

  // @LINE:10
  private lazy val controllers_CategoriesController_updateCategory8_route = Route("PUT",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("categories/"), DynamicPart("id", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_CategoriesController_updateCategory8_invoker = createInvoker(
    CategoriesController_1.get.updateCategory(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CategoriesController",
      "updateCategory",
      Seq(classOf[Int]),
      "PUT",
      this.prefix + """categories/""" + "$" + """id<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:11
  private lazy val controllers_CategoriesController_deleteCategory9_route = Route("DELETE",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("categories/"), DynamicPart("id", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_CategoriesController_deleteCategory9_invoker = createInvoker(
    CategoriesController_1.get.deleteCategory(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CategoriesController",
      "deleteCategory",
      Seq(classOf[Int]),
      "DELETE",
      this.prefix + """categories/""" + "$" + """id<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:13
  private lazy val controllers_CartController_listCartItems10_route = Route("GET",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("cart")))
  )
  private lazy val controllers_CartController_listCartItems10_invoker = createInvoker(
    CartController_2.get.listCartItems,
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CartController",
      "listCartItems",
      Nil,
      "GET",
      this.prefix + """cart""",
      """""",
      Seq()
    )
  )

  // @LINE:14
  private lazy val controllers_CartController_addCartItem11_route = Route("POST",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("cart")))
  )
  private lazy val controllers_CartController_addCartItem11_invoker = createInvoker(
    CartController_2.get.addCartItem,
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CartController",
      "addCartItem",
      Nil,
      "POST",
      this.prefix + """cart""",
      """""",
      Seq()
    )
  )

  // @LINE:15
  private lazy val controllers_CartController_updateCartItem12_route = Route("PUT",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("cart/"), DynamicPart("productId", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_CartController_updateCartItem12_invoker = createInvoker(
    CartController_2.get.updateCartItem(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CartController",
      "updateCartItem",
      Seq(classOf[Int]),
      "PUT",
      this.prefix + """cart/""" + "$" + """productId<[^/]+>""",
      """""",
      Seq()
    )
  )

  // @LINE:16
  private lazy val controllers_CartController_deleteCartItem13_route = Route("DELETE",
    PathPattern(List(StaticPart(this.prefix), StaticPart(this.defaultPrefix), StaticPart("cart/"), DynamicPart("productId", """[^/]+""", encodeable=true)))
  )
  private lazy val controllers_CartController_deleteCartItem13_invoker = createInvoker(
    CartController_2.get.deleteCartItem(fakeValue[Int]),
    play.api.routing.HandlerDef(this.getClass.getClassLoader,
      "router",
      "controllers.CartController",
      "deleteCartItem",
      Seq(classOf[Int]),
      "DELETE",
      this.prefix + """cart/""" + "$" + """productId<[^/]+>""",
      """""",
      Seq()
    )
  )


  def routes: PartialFunction[RequestHeader, Handler] = {
  
    // @LINE:1
    case controllers_ProductsController_listProducts0_route(params@_) =>
      call { 
        controllers_ProductsController_listProducts0_invoker.call(ProductsController_0.get.listProducts)
      }
  
    // @LINE:2
    case controllers_ProductsController_getProduct1_route(params@_) =>
      call(params.fromPath[Int]("id", None)) { (id) =>
        controllers_ProductsController_getProduct1_invoker.call(ProductsController_0.get.getProduct(id))
      }
  
    // @LINE:3
    case controllers_ProductsController_addProduct2_route(params@_) =>
      call { 
        controllers_ProductsController_addProduct2_invoker.call(ProductsController_0.get.addProduct)
      }
  
    // @LINE:4
    case controllers_ProductsController_updateProduct3_route(params@_) =>
      call(params.fromPath[Int]("id", None)) { (id) =>
        controllers_ProductsController_updateProduct3_invoker.call(ProductsController_0.get.updateProduct(id))
      }
  
    // @LINE:5
    case controllers_ProductsController_deleteProduct4_route(params@_) =>
      call(params.fromPath[Int]("id", None)) { (id) =>
        controllers_ProductsController_deleteProduct4_invoker.call(ProductsController_0.get.deleteProduct(id))
      }
  
    // @LINE:7
    case controllers_CategoriesController_listCategories5_route(params@_) =>
      call { 
        controllers_CategoriesController_listCategories5_invoker.call(CategoriesController_1.get.listCategories)
      }
  
    // @LINE:8
    case controllers_CategoriesController_getCategory6_route(params@_) =>
      call(params.fromPath[Int]("id", None)) { (id) =>
        controllers_CategoriesController_getCategory6_invoker.call(CategoriesController_1.get.getCategory(id))
      }
  
    // @LINE:9
    case controllers_CategoriesController_addCategory7_route(params@_) =>
      call { 
        controllers_CategoriesController_addCategory7_invoker.call(CategoriesController_1.get.addCategory)
      }
  
    // @LINE:10
    case controllers_CategoriesController_updateCategory8_route(params@_) =>
      call(params.fromPath[Int]("id", None)) { (id) =>
        controllers_CategoriesController_updateCategory8_invoker.call(CategoriesController_1.get.updateCategory(id))
      }
  
    // @LINE:11
    case controllers_CategoriesController_deleteCategory9_route(params@_) =>
      call(params.fromPath[Int]("id", None)) { (id) =>
        controllers_CategoriesController_deleteCategory9_invoker.call(CategoriesController_1.get.deleteCategory(id))
      }
  
    // @LINE:13
    case controllers_CartController_listCartItems10_route(params@_) =>
      call { 
        controllers_CartController_listCartItems10_invoker.call(CartController_2.get.listCartItems)
      }
  
    // @LINE:14
    case controllers_CartController_addCartItem11_route(params@_) =>
      call { 
        controllers_CartController_addCartItem11_invoker.call(CartController_2.get.addCartItem)
      }
  
    // @LINE:15
    case controllers_CartController_updateCartItem12_route(params@_) =>
      call(params.fromPath[Int]("productId", None)) { (productId) =>
        controllers_CartController_updateCartItem12_invoker.call(CartController_2.get.updateCartItem(productId))
      }
  
    // @LINE:16
    case controllers_CartController_deleteCartItem13_route(params@_) =>
      call(params.fromPath[Int]("productId", None)) { (productId) =>
        controllers_CartController_deleteCartItem13_invoker.call(CartController_2.get.deleteCartItem(productId))
      }
  }
}
