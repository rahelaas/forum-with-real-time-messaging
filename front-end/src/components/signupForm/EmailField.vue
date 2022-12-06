<template>
  <div class="field">
    <div class="ui left icon input big">
      <i class="email icon"></i>

      <my-input
          type="email"
          placeholder="Email"
          autocomplete="off"
          required
          v-model.trim="input"
          @keyup="validateInput"
          @blur="validateInput"
          @input="$emit('update:modelValue', $event.target.value)"
      />

    </div>
    <div class="quickErrors" v-if="errors.email">
      {{ errors.email }}
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";
import MyInput from "@/components/UI/MyInput";

export default {
  name: "EmailField",
  components: {MyInput},
  setup() {
    let input = ref("");
    const { validateEmailField, errors } = useFormValidation();
    const validateInput = () => {
      validateEmailField("email", input.value);
    };


    return { input, errors, validateInput };
  }
}

</script>

