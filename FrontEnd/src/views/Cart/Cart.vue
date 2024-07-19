<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h3 class="pt-3">Keranjang</h3>
      </div>
    </div>
    <!--    loop over all the cart items and display one by one-->
    <div v-for="cartItem in cartItems" :key="cartItem.id" class="row mt-2 pt-3 justify-content-around">
      <div class="col-2"></div>
      <!-- display image -->
      <div class="col-md-3 embed-responsive embed-responsive-16by9">
        <router-link @click.native :to="{ name: 'ShowDetails', params: { id: cartItem.product[0].id } }">
          <img v-bind:src="getImagePathMain(cartItem.product[0].image)" alt="Product Image"
            class="w-100 card-img-top embed-responsive-item" />
        </router-link>
      </div>
      <!-- display product name, quantity and price -->
      <div class="col-md-5 px-3">
        <div class="card-block px-3">
          <h6 class="card-title">
            <router-link @click.native :to="{ name: 'ShowDetails', params: { id: cartItem.product[0].id } }">{{
              cartItem.product[0].name }}
            </router-link>
          </h6>
          <p id="item-price" class="mb-0 font-weight-bold">
            Rp {{ cartItem.product[0].price }} / Produk
          </p>
          <p id="item-quantity" class="mb-0">
            Jumlah : {{ cartItem.quantity }}
          </p>
          <p id="item-total-price" class="mb-0">
            Jumlah Biaya : Rp
            <span class="font-weight-bold">
              {{ cartItem.product[0].price * cartItem.quantity }}</span>
          </p>
          <br /><a href="#" class="text-right"
            @click="deleteItem(cartItem.id, cartItem.product[0].id, cartItem.quantity)">Hapus dari keranjang</a>
        </div>
      </div>
      <div class="col-2"></div>
      <div class="col-12">
        <hr />
      </div>
    </div>

    <!-- display total price -->
    <div class="total-cost pt-2 text-right">
      <h5>Jumlah : {{ totalcost }}</h5>
      <button type="button" class="btn btn-primary confirm" @click="checkout()">
        Konfirmasi Pesanan
      </button>
    </div>
  </div>
</template>

<script>
const axios = require('axios');
export default {
  data() {
    return {
      cartItems: [],
      token: localStorage.getItem('token'),
      totalcost: 0,
      address: [],
    };
  },
  name: 'Cart',
  props: ['baseURL'],
  methods: {
    isDisabled() {
      swal({
        title: "Keranjang Kamu Kosong",
        text: "Silahkan Pilih Produk Terlebih Dahulu",
        icon: "warning",
        buttons: {
          ok: {
            text: "Mulai Belanja!",
            value: "ok",
          }
        },
        closeOnClickOutside: false,
      }).then((value) => {
        if (value === "ok") {
          this.$router.push({ name: 'Home' });
        }
      });
    },
    // fetch all the items in cart
    listCartItems() {
      axios.get(`${this.baseURL}api/carts/users`, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then(response => {
        if (response.data.code === 200) {
          const result = response.data.data;
          // store cartitems and total price in two variables
          this.cartItems = result;
          this.totalcost = 0;
          for (let i = 0; i < this.cartItems.length; i++) {
            for (let j = 0; j < this.cartItems[i].product.length; j++) {
              this.totalcost += this.cartItems[i].product[j].price * this.cartItems[i].quantity;
            }
          }
          if (this.cartItems.length === 0) {
            this.isDisabled();
          }
        }
      })
        .catch((error) => {
          console.log(error);
          this.isDisabled();
        });
    },
    fetchAddress() {
      axios.get(`${this.baseURL}api/address/users`, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        }
      }).then((response) => {
        if (response.data.code === 200 && response.data.data.length > 0) {
          const result = response.data.data[0];
          this.address = result;
        } else {
          this.id = null;
        }
      }).catch((err) => {
        console.log(err);
      })
    },
    checkout() {

      swal({
        title: "Konfirmasi Pesanan",
        text: "Anda yakin ingin melanjutkan pesanan?",
        icon: "warning",
        buttons: ["Cancel", true],
      }).then((willProceed) => {
        if (willProceed) {
          if (!this.address || Object.keys(this.address).length === 0) {
            this.navigateToAddress();
          } else {
            swal({
              title: "Konfirmasi Alamat",
              text: `Alamat Kamu : \n${this.address.nama_penerima} \n${this.address.kabupaten} \n${this.address.kecamatan} \n${this.address.kelurahan} \n${this.address.alamat}\n\n Jika ingin mengubah alamat, klik Ubah Alamat`,
              icon: "warning",
              buttons: {
                cancel: "Batalkan",
                update: {
                  text: "Ubah Alamat",
                  value: "update",
                },
                ok: true,
              }
            }).then((value) => {
              if (value === "update") {
                this.navigateToAddress();
              } else if (value) {
                localStorage.setItem('totalcost', this.totalcost);
                axios.post(`${this.baseURL}api/orders`, { cart: this.cartItems.id }, { headers: { Authorization: `Bearer ${this.token}` } })
                  .then((response) => {
                    if (response.data.code === 200) {
                      swal("Thanks For Orders", "", "success");
                      this.$router.push({ name: 'Checkout', params: { orderId: response.data.data.id, totalcost: this.totalcost } });
                    } else {
                      swal("Error", response.data.message, "error");
                    }
                  }).catch((error) => {
                    console.log(error);
                    swal("Error", "An error occurred while placing the order", "error");
                  });
              }
            })
          }
        }
      });
    },
    navigateToAddress() {
      window.location.href = this.$router.resolve({ name: 'Address' }).href;
    },
    deleteItem(itemId, productId, quantityItem) {
      itemId = parseInt(itemId, 10);
      const config = {
        headers: {
          Authorization: `Bearer ${this.token}`
        },
        data: {
          id: itemId,
          product_id: productId,
          quantity: quantityItem
        }
      };
      axios
        .delete(`${this.baseURL}api/remove/carts`, config)
        .then((response) => {
          if (response.data.code === 200) {
            swal({
              text: "Produk berhasil dihapus dari keranjang!",
              icon: "success",
              closeOnClickOutside: false,
            });
            const cartIndex = this.cartItems.findIndex(item => item.id === itemId);
            const deletedItemCost = this.cartItems[cartIndex].product.reduce((acc, product) => acc + (product.price * this.cartItems[cartIndex].quantity), 0);

            this.totalcost -= deletedItemCost;

            this.cartItems = this.cartItems.filter(item => item.id !== itemId);
          }
          this.$emit('fetchData');
        })
        .catch((error) => {
          console.log(error);
          // Handle the error, e.g., show an error message
        });
    },
    showDetails(productId) {
      this.$router.push({
        name: 'ShowDetails',
        params: { id: productId },
      });
    },
    getImagePathMain(image) {
      try {
        return require(`../../assets/img-main/${image}`);
      } catch (e) {
        return '';
      }
    },
  },
  mounted() {
    this.token = localStorage.getItem('token');
    this.listCartItems();
    this.fetchAddress();
    //this.checkout();
  },
};
</script>

<style scoped>
h4,
h5 {
  font-family: 'Roboto', sans-serif;
  color: #484848;
  font-weight: 700;
}

.embed-responsive .card-img-top {
  object-fit: cover;
}

#item-price {
  color: #232f3e;
}

#item-quantity {
  float: left;
}

#item-total-price {
  float: right;
}

.confirm {
  font-weight: bold;
  font-size: larger;
}
</style>
