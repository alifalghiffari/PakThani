<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">Daftar Transaksi</h4>
      </div>
    </div>
    <!--        for each order display -->
    <div v-for="order in orderList" :key="order.id" class="row mt-2 pt-3 justify-content-around">
      <!-- <div class="col-2"></div> -->
      <!-- <div class="col-md-3 embed-responsive embed-responsive-16by9"> -->
        <!--                display image in left-->
        <!-- <img v-bind:src="order.imageURL" class="w-100 card-img-top embed-responsive-item"> -->
      <!-- </div> -->
      <div class="col-md-5 px-3">
        <div class="card-block px-3">
          <h6 class="card-title">
            <router-link v-bind:to="'/order/'+order.id">Nomor Pesanan : {{order.id}}</router-link>
          </h6>
          <p class="mb-0">{{order.totalItems}} Item<span v-if="order.totalItems > 1"></span></p>
          <p id="item-price" class="mb-0 font-weight-bold">Jumlah Biaya : Rp {{order.totalPrice}}</p>
          <!-- <p id="item-username">Ordered By : {{order.username}}</p> -->
          <p id="item-status">Status : {{ order.orderStatus }}</p>
        </div>
      </div>
      <div class="col-2"></div>
      <div class="col-12"><hr></div>
    </div>
  </div>

</template>

<script>
  const axios = require('axios')
  export default {

    data() {
      return {
        token: null,
        orderList : []
      }
    },
    props:["baseURL"],
    name: 'Order',
    methods: {
      fecthOrder() {
        axios.get(`${this.baseURL}api/orders/users`, {
          headers: {
              Authorization: `Bearer ${this.token}`,
          },
        }).then((response) => {
          if (response.status === 200) {
            const orders = response.data.data;
            orders.forEach((order) => {
              this.orderList.push({
                  id: order.id,
                  totalItems: order.total_items,
                  totalPrice: order.total_price,
                  orderStatus: order.order_status,
                  paymentStatus: order.payment_status,
              });
            });
          }
        }).catch((err) => {
            console.error(err);
        });
      }
      // list of order histories
      // listOrders(){
      //   axios.get(`${this.baseURL}api/orders/users`, {
      //     headers: {
      //       Authorization: `Bearer ${this.token}`,
      //     },
      //   })
      //     .then((response) => {
      //         if(response.status==200){
      //           this.orders = response.data.data;
                // for each order populate orderList
                // this.orders.forEach((order) => {
                //   this.orderList.push({
                //     id: order.id,
                //     totalCost: order.totalPrice,
                //     orderStatus: order.order_status,
                    // get short date
                    //orderdate: order.createdDate.substring(0,10),
                    // get image of the first orderItem of the order
                    //imageURL: order.orderItems[0].product.imageURL,
                    // get total items
                    //totalItems: order.orderItems.length
            //       })
            //     })
            //   }
            // },
            // (error)=>{
            //   console.log(error)
            // });
      //},
    },
    mounted() {
      this.token = localStorage.getItem("token");
      //this.listOrders();
      this.fecthOrder();
    },
  };

</script>

<style scoped>
  h4, h5 {
    font-family: 'Roboto', sans-serif;
    color: #484848;
    font-weight: 700;
  }

  .embed-responsive .card-img-top {
    object-fit: cover;
  }
</style>
