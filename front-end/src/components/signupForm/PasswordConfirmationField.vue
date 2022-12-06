<template>
  <div class="field">
    <div class="ui left icon input big">
      <i class="lock icon"></i>

      <my-input
          v-model.trim="input"
          type="password"
          placeholder="Re-enter your password"
          autocomplete="off"
          required
          @keyup="validateInput"
          @blur="validateInput"
          @input="$emit('update:modelValue', $event.target.value)"
      />

    </div>
    <div class="quickErrors" v-if="errors.confirmation">
      {{ errors.confirmation }}
    </div>
  </div>
</template>

<script>
import {ref} from "vue";
import useFormValidation from "@/modules/useFormValidation";
import MyInput from "@/components/UI/MyInput";

export default {
  name: "PasswordConfirmationField",
  components: {MyInput},

  setup() {
    let input = ref(null);
    const {validateConfirmPasswordField, errors} = useFormValidation();
    const validateInput = () => {
      validateConfirmPasswordField("confirmation", input.value);
    };
    return {input, errors, validateInput};
  },
};
</script>

