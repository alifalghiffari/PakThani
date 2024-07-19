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
            <h2 class="pt-4 pl-4">Verifikasi Email</h2>
            <form @submit="verify" class="pt-4 pl-4 pr-4">
              <div class="form-group">
                <label>Kata Sandi Baru</label>
                <input type="password" name="password" id="password" class="form-control" v-model="password" required
                :class="{ 'is-invalid': password && !passwordRegex.test(password) }" />
                <div class="invalid-feedback" v-if="password && !passwordRegex.test(password)">
                    Kata sandi harus terdiri dari minimal 8 karakter, terdiri dari satu huruf besar, satu huruf kecil, satu
                    angka, dan satu karakter khusus.
                </div>
              </div>
              <div class="form-group">
                <label>Konfirmasi Sandi</label>
                <input type="password" class="form-control" v-model="passwordConfirm" required />
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
        name: "ResetPassword",
        data() {
            return {
            password: null,
            passwordConfirm: null,
            passwordRegex: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/,
            token: localStorage.getItem('token'),
            };
        },
        methods: {
            async verify(e) {
                e.preventDefault();
          
                if (this.password === this.passwordConfirm && this.passwordRegex.test(this.password)) {
                    const user = {
                        new_password: this.password,
                    };
                    await axios
                    .post(`http://localhost:3000/api/users/reset-password`, user, {
                        headers: {
                            Authorization: `Bearer ${this.token}`,
                        },
                    })
                    .then((res) => {
                        if (res.data.code === 200) {
                            swal({
                                text: "Sandi telah di ubah!",
                                icon: "success",
                                closeOnClickOutside: false,
                            });
                            this.$router.push({ name: "Signin" });
                            localStorage.removeItem('token');
                        }
                    })
                    .catch((err) => {
                        console.log(err);
                    });
                } else {
                    
                    swal({
                    text: "Password tidak sesuai atau tidak memenuhi syarat!",
                    icon: "error",
                    closeOnClickOutside: false,
                    });
                }
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
    
    