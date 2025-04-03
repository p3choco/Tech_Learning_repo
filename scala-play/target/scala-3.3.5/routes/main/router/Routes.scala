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
  val prefix: String
) extends GeneratedRouter {

  @javax.inject.Inject()
  def this(errorHandler: play.api.http.HttpErrorHandler,
    // @LINE:1
    ProductsController_0: javax.inject.Provider[controllers.ProductsController]
  ) = this(errorHandler, ProductsController_0, "/")

  def withPrefix(addPrefix: String): Routes = {
    val prefix = play.api.routing.Router.concatPrefix(addPrefix, this.prefix)
    router.RoutesPrefix.setPrefix(prefix)
    new Routes(errorHandler, ProductsController_0, prefix)
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
  }
}
