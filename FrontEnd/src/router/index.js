import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import AboutUs from '../views/AboutUs.vue'
import ContactUs from '../views/ContactUs.vue'
import OurBusiness from '../views/OurBusiness.vue'

import Admin from '../views/Admin/Admin.vue'
import Gallery from '../views/Admin/Gallery.vue'
import AddImage from '../views/Admin/AddImage.vue'

import PageNotFound from '../views/PageNotFound.vue'

import Product from '../views/Product/Product.vue'
import AddProduct from '../views/Product/AddProduct.vue'
import EditProduct from '../views/Product/EditProduct.vue'
import ShowDetails from '../views/Product/ShowDetails.vue'
import Wishlist from '../views/Product/Wishlist.vue'
import Cart from '../views/Cart/Cart.vue'
import Checkout from '../views/Checkout/Checkout.vue'
import Address from '../views/Checkout/Address.vue'
import Order from '../views/Orders/Order.vue'
import OrderAdmin from '../views/Orders/OrderAdmin.vue'
import OrderDetailsAdmin from '../views/Orders/OrderDetailsAdmin.vue'

import Category from '../views/Category/Category.vue'
import AddCategory from '../views/Category/AddCategory.vue'
import EditCategory from '../views/Category/EditCategory.vue'
import ListProducts from '../views/Category/ListProducts.vue'
import Signup from '../views/Signup.vue'
import Signin from '../views/Signin.vue'
import VerifyAccount from '../views/VerifyAccount.vue'
import VerifyEmail from '../views/VerifyEmail.vue'
import VerifyOTP from '../views/VerifyOTP.vue'
import ResetPassword from '../views/ResetPassword.vue'

import Success from '../helper/payment/Success.vue'
import Failed from '../helper/payment/Failed.vue'

import OrderDetails from "../views/Orders/OrderDetails";



const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/about',
    name: 'AboutUs',
    component: AboutUs
  },
  {
    path: '/about/contact-us',
    name: 'ContactUs',
    component: ContactUs
  },
  {
    path: '/about/our-business',
    name: 'OurBusiness',
    component: OurBusiness
  },
  //Admin routes
  {
    path: '/admin',
    name: 'Admin',
    component: Admin
  },
  {
    path : '/admin/gallery',
    name : 'Gallery',
    component : Gallery
  },
  {
    path : '/admin/gallery/add',
    name : 'AddImage',
    component : AddImage
  },
  //Product routes
  {
    path: '/product',
    name: 'Product',
    component: Product
  },
  {
    path: '/admin/product',
    name: 'AdminProduct',
    component: Product
  },
  {
    path: '/admin/product/add',
    name: 'AddProduct',
    component: AddProduct
  },
  {
    path: '/admin/product/:id',
    name: 'EditProduct',
    component: EditProduct,
  },
  {
    path : '/product/show/:id',
    name : 'ShowDetails',
    component: ShowDetails
  },
  //Category routes
  {
    path: '/category',
    name: 'Category',
    component: Category
  },
  {
    path: '/admin/category',
    name: 'AdminCategory',
    component: Category
  },
  {
    path: '/admin/category/add',
    name: 'AddCategory',
    component: AddCategory
  },
  {
    path: '/admin/category/:id',
    name: 'EditCategory',
    component: EditCategory
  },
  {
    path : '/category/show/:id',
    name : 'ListProducts',
    component: ListProducts
  },
  {
    path: '/admin/order/',
    name: 'OrderAdmin',
    component: OrderAdmin
  },
  {
    path: '/admin/order/:id',
    name: 'OrderDetailsAdmin',
    component: OrderDetailsAdmin
  },
  //Page Not found
  {
    path : '/:catchAll(.*)',
    name : 'PageNotFound',
    component : PageNotFound
  },
  //Signin and Signup
  {
    path: '/signup',
    name: 'Signup',
    component: Signup
  },
  {
    path: '/signin',
    name: 'Signin',
    component: Signin
  },
  {
    path: '/verify-account',
    name: 'VerifyAccount',
    component: VerifyAccount
  },
  {
    path: '/verify-email',
    name: 'Verify',
    component: VerifyEmail
  },
  {
    path: '/verify-otp',
    name: 'VerifyOTP',
    component: VerifyOTP
  },
  {
    path: '/reset-password',
    name: 'ResetPassword',
    component: ResetPassword
  },
  {
    path: '/wishlist',
    name: 'Wishlist',
    component: Wishlist
  },
  {
    path : '/cart',
    name : 'Cart',
    component : Cart
  },
  {
    path : '/checkout',
    name : 'Checkout',
    component : Checkout
  },
  {
    path : '/address',
    name : 'Address',
    component : Address
  },
  {
    path : '/order',
    name : 'Order',
    component : Order
  },
  {
    path: '/payment/success',
    name: 'PaymentSuccess',
    component:Success
  },
  {
    path: '/payment/failed',
    name: 'FailedPayment',
    component:Failed
  },
  {
    path:'/order/:id',
    name:'OrderDetails',
    component: OrderDetails
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const requiresAuth = ['Cart', 'Order', 'Address', 'Admin', 'AddProduct', 'EditProduct', 'AddCategory', 'EditCategory', 'AddImage', 'OrderAdmin', 'OrderDetailsAdmin', 'ResetPassword'];
  const authRequired = requiresAuth.includes(to.name);

  const isLoggedIn = localStorage.getItem('token');
  const user = localStorage.getItem('user');

  if (authRequired) {
    if (!isLoggedIn) {
      // Redirect ke halaman login jika tidak ada token
      next('/signin');
    } else {
      // Periksa peran pengguna di sini
      if (user === "admin") {
        //how to create this user can access all page
        next();
      } else {
        //how to create this user can access all page except admin page
        if (to.name === 'Admin' || to.name === 'AddProduct' || to.name === 'EditProduct' || to.name === 'AddCategory' || to.name === 'EditCategory' || to.name === 'AddImage' || to.name === 'OrderAdmin' || to.name === 'OrderDetailsAdmin') {
          next('/');
        } else {
          next();
        }
      }
    }
  } else {
    // Halaman yang tidak memerlukan otorisasi khusus
    next();
  }
});


export default router
