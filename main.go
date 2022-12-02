/**
@apiDefine admin User access only
 This optional description belong to to the group admin.
*/

/**
@apiDefine MyErrors
@apiError 429 Too Many Requests
@apiErrorExample (429) Error-Response:
        HTTP/1.1 429 OK
        {
          "error": "The user has sent too many requests in a given amount of time (“rate limiting”)."
        }
@apiError 503 Service Unavailable
@apiErrorExample (503) Error-Response:
      HTTP/1.1 503 OK
      {
        "error": "The server is not ready to handle the request"
      }
*/

/**
@api {get} /products/ All products list
@apiVersion 0.1.0
@apiGroup Products
@apiName List of products
@apiPermission All
@apiSuccess (200) {object} products List of products
@apiSuccessExample {object} Success-Response:
      HTTP/1.1 200 OK
      {
		"products": [{
			"id": "123"
			"name": "Banana"
			"description": "Green banana"
			"ammount": "2987"
			"price": "2$"
			"category": "Fruit"
		}]
      }

@apiUse MyErrors
*/

/**

@api {get} /products/search/:name Search for a product
@apiName Product
@apiGroup Products
@apiVersion 0.1.0
@apiPermission All
@apiParam {Number} name Name of product
@apiParam {Number} [id] Id of product
@apiSuccess {object} products Product found
@apiSuccessExample {object} Success-Response:
      HTTP/1.1 200 OK
      {
	"products": [{
                "id": "123"
                "name": "Banana"
                "description": "Green banana"
                "ammount": "2987"
                "price": "2$"
                "category": "Fruit"
        }]

      }

@apiError 404 Not found
@apiErrorExample (404) Error-Response:
	HTTP/1.1 404 OK
	{
	  "error": "We do not have that product"
	}
@apiUse MyErrors
*/

/**

@api {put} /products/modify/:id Product modification
@apiName Product modification
@apiPermission admin
@apiVersion 0.1.0
@apiGroup Products
@apiParam {number} id Product id
@apiParam {string} [name] Product name
@apiParam {string} [description] Product description
@apiParam {string} [ammount] Ammount of product left
@apiParam {string} [price] Price of product
@apiParam {string} [category] Product category
@apiHeader {String} TOKEN Users unique token.
@apiSampleRequest https://api.shop.com/v1
@apiSuccess (200) {object} products Modified product
@apiSuccessExample {object} Success-Response:
      HTTP/1.1 200 OK
      {
        "products": [{
                "id": "123"
                "name": "Banana"
                "description": "Green banana"
                "ammount": "2987"
                "price": "2$"
                "category": "Fruit"
        }]
      }
@apiError Unauthorized
@apiErrorExample (401) Error-Response:
	HTTP/1.1 401 OK
	{
	  "error": "Unable to authenticate you."
	}

@apiUse MyErrors
*/

/**

@api {post} /basket/:id/:basket Add product to basket
@apiName Add product to basket
@apiPermission All
@apiGroup Basket
@apiVersion 0.1.0
@apiParam {number} id Product id
@apiParam {number} basket Basket id
@apiSuccess {object} id  Product id
@apiSuccess {object} basket  Basket id
@apiSuccess {object} name  Product name
@apiSuccess {object} ammount  Product ammount
@apiSuccessExample {object} Success-Response:
      HTTP/1.1 200 OK
      {
		"id": "123"
		"basket": "321"
		"name": "Apple"
		"ammount": "67"
      }

@apiUse MyErrors
*/
