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
          <h2 class="pt-4 pl-4">Masuk</h2>
          <form @submit="signin" class="pt-4 pl-4 pr-4">
            <div class="form-group">
              <label>Email</label>
              <input type="email" class="form-control" v-model="email" required />
            </div>
            <div class="form-group">
              <label>Password</label>
              <input type="password" class="form-control" v-model="password" required />
            </div>
            <router-link :to="{ name: 'Verify' }" class=""><small class="text-muted text-start">Lupa Kata Sandi?</small></router-link>
            <div class="d-flex justify-content-center mt-2">
              <button type="submit" class="btn btn-primary py-0">
                Lanjutkan
                <div v-if="loading" class="spinner-border spinner-border-sm" role="status">
                  <span class="sr-only">Loading...</span>
                </div>
              </button>
            </div>
            <div class="d-flex align-items-center my-3">
              <div class="lineOr"></div>
              <span class="p-0 m-0 text-center middleOr">atau</span>
              <div class="lineOr"></div>
            </div>
            <a href="javascript:void(0);" @click="signInWithGoogle()" class="d-flex justify-content-center google-signin-btn mb-2">
              <svg viewBox="0 0 24 24" width="24" height="24" xmlns="http://www.w3.org/2000/svg">
                <g transform="matrix(1, 0, 0, 1, 27.009001, -39.238998)">
                  <path fill="#4285F4" d="M -3.264 51.509 C -3.264 50.719 -3.334 49.969 -3.454 49.239 L -14.754 49.239 L -14.754 53.749 L -8.284 53.749 C -8.574 55.229 -9.424 56.479 -10.684 57.329 L -10.684 60.329 L -6.824 60.329 C -4.564 58.239 -3.264 55.159 -3.264 51.509 Z"/>
                  <path fill="#34A853" d="M -14.754 63.239 C -11.514 63.239 -8.804 62.159 -6.824 60.329 L -10.684 57.329 C -11.764 58.049 -13.134 58.489 -14.754 58.489 C -17.884 58.489 -20.534 56.379 -21.484 53.529 L -25.464 53.529 L -25.464 56.619 C -23.494 60.539 -19.444 63.239 -14.754 63.239 Z"/>
                  <path fill="#FBBC05" d="M -21.484 53.529 C -21.734 52.809 -21.864 52.039 -21.864 51.239 C -21.864 50.439 -21.724 49.669 -21.484 48.949 L -21.484 45.859 L -25.464 45.859 C -26.284 47.479 -26.754 49.299 -26.754 51.239 C -26.754 53.179 -26.284 54.999 -25.464 56.619 L -21.484 53.529 Z"/>
                  <path fill="#EA4335" d="M -14.754 43.989 C -12.984 43.989 -11.404 44.599 -10.154 45.789 L -6.734 42.369 C -8.804 40.429 -11.514 39.239 -14.754 39.239 C -19.444 39.239 -23.494 41.939 -25.464 45.859 L -21.484 48.949 C -20.534 46.099 -17.884 43.989 -14.754 43.989 Z"/>
                </g>
              </svg>
              <span>Masuk dengan Google</span>
            </a>
          </form>
          <hr />
          <small class="form-text text-muted pt-2 text-center">Yuk daftar PakThani</small>
          <p class="text-center">
            <router-link :to="{ name: 'Signup' }" class="btn btn-dark text-center mx-auto px-5 py-1 mb-2">Buat Akun</router-link>
          </p>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import { googleSdkLoaded } from "vue3-google-login";
import axios from 'axios';

export default {
  name: "Signin",
  data() {
    return {
      email: null,
      password: null,
      loading: false,
    };
  },
  methods: {
    async signin(e) {
      e.preventDefault();
      this.loading = true;

      const user = {
        email: this.email,
        password: this.password,
      };

      try {
        const res = await axios.post(`http://localhost:3000/api/users/login`, user);
        if (res.data.code === 500) {
          swal({
            text: "Email atau Password Salah!",
            icon: "error",
            closeOnClickOutside: false,
          });
        } else if (res.data.code === 200) {
          localStorage.setItem("token", res.data.data.token);
          localStorage.setItem("user", res.data.data.role);
          this.$emit("fetchData");
          this.$router.push({ name: "Home" });
        }
      } catch (err) {
        swal({
          text: "Tidak bisa Masuk! Coba Lagi!",
          icon: "error",
          closeOnClickOutside: false,
        });
        console.log(err);
      } finally {
        this.loading = false;
      }
    },
    signInWithGoogle() {
      googleSdkLoaded(google => {
        google.accounts.oauth2
          .initCodeClient({
            client_id:
              "1083611601525-sm55jlhpioajjvbp811ikq5bbn2m04vh.apps.googleusercontent.com",
            scope: "email profile openid",
            redirect_uri: "http://localhost:8080/auth/callback",
            callback: response => {
              if (response.code) {
                console.log(response.code);
                this.sendCodeToBackend(response.code);
              }
            }
          })
          .requestCode();
      });
    },
    async sendCodeToBackend(code) {
      try {
        const response = await axios.get(`http://localhost:3000/api/users/login-google?state=${code}`);
        if (response.data.code === 200) {
          localStorage.setItem('token', response.data.data.token);
          localStorage.setItem('user', response.data.data.role);
          this.$emit('fetchData');
          this.$router.push({ name: 'Home' });
        } else {
          swal({
            text: "Google Sign-In failed!",
            icon: "error",
            closeOnClickOutside: false,
          });
        }
      } catch (error) {
        console.error("Failed to send ID token to backend:", error);
      }
    }
  },
  mounted() {
    this.loading = false;
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

.google-signin-btn {
  justify-content: center;
  text-align: center;
  display: inline-flex;
  align-items: center;
  padding: 10px 16px;
  background-color: #fff;
  color: #757575;
  border-radius: 4px;
  font-family: 'Roboto', sans-serif;
  font-size: 14px;
  font-weight: 500;
  text-decoration: none;
  transition: background-color 0.3s;
}

.google-signin-btn:hover {
  background-color: #f7f7f7;
}

.google-signin-btn svg {
  margin-right: 10px;
}

.lineOr {
  background-color: #dbdbdb;
  flex: 1;
  height: 1px;
  width: 100%;
}

.middleOr {
  color: #ccc;
  font-size: .75rem;
  padding: 0 16px;
  text-transform: uppercase;
}
</style>

