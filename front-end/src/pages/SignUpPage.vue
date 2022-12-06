<template>

  <form @click.prevent @submit.prevent>

      <h3>Please Register</h3>


      <label for="username">Username</label>
      <UsernameField v-model="usere.username" id="username" />

      <label for="firstName">First name</label>
      <FirstNameField v-model="usere.firstname" id="firstName"/>

      <label for="lastName">Last name</label>
      <LastNameField v-model="usere.lastname" id="lastName" />

      <label for="email">Email</label>
      <EmailField v-model="usere.email" id="email"/>

      <GenderField v-model="usere.gender" :label="genderLabel" />

      <BirthYearField id="years" v-model="usere.year" :label="yearLabel" />

      <label for="password">Password</label>
      <PasswordField v-model.trim="usere.password" id="password" :password="sendPassword"/>

     <!-- <label for="confirm">Password confirmation</label>
      <PasswordConfirmationField v-model.trim="usere.confirmation" ref="confirmation"  id="confirm"/> -->

    <p v-if="errorse.length">
      <b>Please correct the following error(s):</b>
    <ul>
      <li :key="error" class="error" v-for="error in errorse">{{ error }}</li>
    </ul>
    </p>

    <div class="buttons">

      <my-button class="btn" type="button" @click="$router.push('/')">Cancel</my-button>
      <my-button class="btn" type="button" :disabled="isSignupButtonDisabled" @click="submitRegistration">Submit</my-button>

    </div>
  </form>

  <br /><br />
  <div>
    <strong>Already Registered Member? </strong>
    <router-link to="/">Log In</router-link> here
  </div>

</template>

<script>
import MyButton from "@/components/UI/MyButton";
import axios from "axios";
import {reactive} from "@vue/reactivity";
import useFormValidation from "@/modules/useFormValidation";
import useSubmitButtonState from "@/modules/UseSubmitButtonState";
import UsernameField from "@/components/signupForm/UsernameField";
import FirstNameField from "@/components/signupForm/FirstNameField";
import LastNameField from "@/components/signupForm/LastNameField";
import EmailField from "@/components/signupForm/EmailField";
import GenderField from "@/components/signupForm/GenderField";
import PasswordField from "@/components/signupForm/PasswordField";
// import PasswordConfirmationField from "@/components/signupForm/PasswordConfirmationField";
import BirthYearField from "@/components/signupForm/BirthYearField";
//axios.defaults.withCredentials = true;

export default {
  name: "SignUp",
  components: {
    MyButton,
    UsernameField,
    FirstNameField,
    LastNameField,
    EmailField,
    GenderField,
    BirthYearField,
    PasswordField,
   // PasswordConfirmationField
  },

  setup() {
    let usere = reactive({
      username: '',
      firstname: '',
      lastname: '',
      email: '',
      gender: '',
      year: '',
      password: '',
     // confirmation: '',


    });

    const { errors } = useFormValidation();
    const { isSignupButtonDisabled } = useSubmitButtonState(usere, errors);
    return { usere, isSignupButtonDisabled}
  },

  data() {
    return {

      genderLabel: 'Gender',
      yearLabel: 'Year of birth',
      sendPassword: 'usere.password',
      socket: null,
      info: null,
      errorse: [],
      msg: ""
    };
  },

  methods: {
    /*  checkForm: function () {
     /* this.errorse = [];
        if (this.usere.password !== this.usere.confirmation) {
          this.errorse.push('Password and password confirmation don\'t match.');
          this.usere.password = ''
          this.usere.confirmation = ''
          this.resetInput(this.usere.password)
          this.resetInput(this.usere.confirmation)
        }
    },
  */
    submitRegistration() {
      //this.checkForm(this.usere)
      if (this.errorse.length < 1){
        axios({
          method: "post",
          url: "http://localhost:9200/signup",
          data: this.usere,
          headers: {'Content-Type': 'text/plain'},
          withCredentials: true,
          origin: true,
        }).then((result) => {
          console.log("result ", result)
          this.msg = result.data["msg"];
          this.flag = result.data["check"];
          if (this.flag === true) {
            alert(result.data["msg"]);
            setTimeout(() => {this.$router.push({ path: "/" })
            }, 2000);
          } else {
            alert(result.data["msg"]);
            if (result.data["msg"]==="Username and email are already taken."){
              this.usere.username = ""
              this.usere.email = ""
              location.reload()
            }
            if (result.data["msg"]==="Email is taken."){
              this.usere.email = ""
              location.reload()
            }
            if (result.data["msg"]==="Username is taken."){
              this.usere.username = ""
              location.reload()
            }
          }

        }).catch((error) => {
          if (error.response) {
            console.log(error.response);
          } else if (error.request) {
            console.log("network error");
          } else {
            console.log(error);
          }
        });
      }
    },
  },
};
</script>

<style scoped>
form {
  margin-left: 10%;
  margin-top: 2%;
  display: flex;
  flex-direction: column;
  width: 500px;
}
.error {
  color: darkred;
}
.quickErrors{
  color: red;
}

</style>

<style>
select {
  width: 50%;
  border: 1px solid teal;
  padding: 10px 15px;
}

select:invalid {
  color: gray;
  font-family: KindelSerif, serif;
}
label, .label {
  margin-top: 10px;
  margin-bottom: 4px;
  font-family: KindelSerif, serif ;
}
</style>