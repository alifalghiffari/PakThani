<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <router-link class="navbar-brand" :to="{ name: 'Home' }">
                <img id="logo" src="../assets/logo.png" />
            </router-link>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ml-auto text-white mt-2 mt-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link" :to="{ name: 'AboutUs' }">Beranda</router-link>
                    </li>
                    <!-- <li class="nav-item ml-4">
                        <router-link class="nav-link" :to="{ name: 'OurBusiness' }">Bisnis Kami</router-link>
                    </li> -->
                    <li class="nav-item ml-4">
                        <router-link class="nav-link" :to="{ name: 'ContactUs' }">Kontak</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

    <div class="page-img bg-cover" id="bg10">
        <div class="container">
            <div class="row align-items-center justify-content-center" style="min-height: 70vh;">
                <div class="col-md-9 text-center">
                    <p class="breadcrumbs text-light">
                        <span class="mr-2">
                            <router-link class="text-light" :to="{ name: 'Home' }">Home</router-link>
                        </span>
                        <span>Contact Us</span>
                    </p>
                    <h1 class="mb-0 bread text-light">Contact Us</h1>
                </div>
            </div>
        </div>
    </div>

    <div class="container-fluid">
        <div class="row justify-content-center align-items-center">
            <div class="col-md-5 " id="bg12">
                
            </div>
            <div class="col-md-5 form-column">
                <form action="#" class="contact-form">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Your Name">
                </div>
                <div class="form-group">
                    <input type="email" class="form-control" placeholder="Your Email">
                </div>
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Subject">
                </div>
                <div class="form-group">
                    <textarea name="message" cols="30" rows="7" class="form-control" placeholder="Message"></textarea>
                </div>
                <div class="form-group">
                    <button type="submit" class="btn btn-success btn-block">Send Message</button>
                </div>
                </form>
            </div>
        </div>
    </div>

</template>

<script>
export default {
    name: 'ContactUs',
};
</script>

<style scoped>
.page-img {
    min-height: 70vh;
}

.bg-cover {
    background-size: cover !important;
}

#bg10 {
    background: url('../assets/vegetables.jpg') no-repeat center center;
    background-size: cover;
    height: 70vh;
}

#bg12 {
    background: url('../assets/vegetables-fruits.jpg') no-repeat center center;
    background-size: cover;
    height: 100vh;
}

.image-column, .form-column {
  padding: 1px;
}
.contact-form {
  background-color: #fff;
  padding: 30px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0,0,0,0.1);
}
.btn-success {
  background-color: #4CAF50;
  border: none;
}
.form-control {
  margin-bottom: 15px;
}

</style>
