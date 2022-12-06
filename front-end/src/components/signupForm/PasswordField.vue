<template>
  <div class="field">
    <div class="ui left icon input big">

      <my-input
          type="password"
          placeholder="Create a password"
          autocomplete="off"
          required
          v-model.trim="input"
          @keyup="validateInput"
          @blur="validateInput"
          @input="$emit('update:modelValue', $event.target.value)"
      />

    </div>
    <div class="quickErrors" v-if="errors.password">
      {{ errors.password }}
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";
import MyInput from "@/components/UI/MyInput";

export default {
  name: "PasswordField",
  components:{MyInput},

  setup() {
    let input = ref("");
    const { validatePasswordField, errors } = useFormValidation();
    const validateInput = () => {
      validatePasswordField("password", input.value);
    };
    return { input, errors, validateInput };
  },
};
</script>

