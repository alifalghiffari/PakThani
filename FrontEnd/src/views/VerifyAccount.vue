<template>
  <div class="container">
    <!-- Logo Div -->
    <div class="row">
      <div class="col-12 text-center pt-3">
        <router-link :to="{ name: 'Home' }">
          <img id="logo" src="../assets/logo.png" />
        </router-link>
      </div>
    </div>

    <div class="row">
      <div class="col-12 justify-content-center d-flex flex-row pt-3">
        <div id="signin-div" class="flex-item border">
          <h2 class="pt-4 pl-4">Verifikasi Akun</h2>
          <form @submit="verify" class="pt-4 pl-4 pr-4">
            <div class="form-group">
              <label>Email</label>
              <input type="email" class="form-control" v-model="email" required />
            </div>
            <div class="form-group">
              <label>OTP</label>
              <input type="text" class="form-control" v-model="otp" required />
            </div>
            <div class="d-flex justify-content-center mt-2 pb-4">
              <button type="submit" class="btn btn-primary py-0">
                Lanjutkan
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
export default {
  name: "VerifyAccount",
  data() {
    return {
      email: null,
      otp: null,
    };
  },
  methods: {
    async verify(e) {
      e.preventDefault();

      const user = {
        email: this.email,
        verification_token: this.otp
      };
      console.log(user);
      await axios
        .post(`http://localhost:3000/api/users/verify-email`, user)
        .then((res) => {
          if (res.data.code === 200 && res.data.data.id === 0) {
            swal({
              text: "Kode OTP atau Email salah!",
              icon: "error",
              closeOnClickOutside: false,
            });
          } else {
            swal({
              text: "Akun berhasil di verifikasi!",
              icon: "success",
              closeOnClickOutside: false,
            });
            this.$router.push({ name: "Signin" });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    }
  },
  mounted() {

  },
};
</script>

<style scoped>
.btn-dark {
  background-color: #e7e9ec;
  color: #000;
  font-size: smaller;
  border-radius: 0;
  border-color: #adb1b8 #a2a6ac #a2a6ac;
}

.btn-primary {
  background-color: #f0c14b;
  color: black;
  border-color: #a88734 #9c7e31 #846a29;
  border-radius: 0;
}

#logo {
  width: 150px;
}

@media only screen and (min-width: 992px) {
  #signin-div {
    width: 40%;
  }
}
</style>