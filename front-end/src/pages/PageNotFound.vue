<template>
  <div class="parent">
    <section class="page_404">
      <div class="container">
        <div class="row">
          <div class="col-sm-12 ">
            <div class="col-sm-10 col-sm-offset-1  text-center">
              <div class="four_zero_four_bg">
                <h1 class="text-center ">404</h1>


              </div>

              <div class="contant_box_404">
                <h3 class="h2">
                  Looks like you're lost
                </h3>

                <p>The page you are looking for not available!
                  <br>Let's go Home and try from there!
                </p>

                <a v-if="userStatus === false" href="http://localhost:8080/" class="link_404">Go to Home</a>
                <a v-if="userStatus === true" href="http://localhost:8080/posts" class="link_404">Go back to Posts</a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import store from "@/store";
export default {
  computed: {
    userStatus() {
      if (store.state.user.isLoggedIn === false) {
        return false
      } else {
        return true
      }
    },
  },
}
</script>

<style scoped>
.page_404{
  padding:40px 0;
  background:#fff;
  font-family: 'Arvo', serif;
  justify-content: center;
  display: flex;

}

.page_404  img{
  width:100%;
}

.four_zero_four_bg{

  background-image: url(https://cdn.dribbble.com/users/285475/screenshots/2083086/dribbble_1.gif);
  height: 400px;
  background-position: center;
}


.four_zero_four_bg h1{
  font-size:80px;
}

.four_zero_four_bg h3{
  font-size:80px;
}

.link_404{
  color: #fff!important;
  padding: 10px 20px;
  background: #39ac31;
  margin: 20px 0;
  display: inline-block;}
.contant_box_404{ margin-top:-50px;}



</style>