<template>
  <div class="field">
    <div class="ui left icon input big">
      <i class="user icon"></i>

      <my-input
          type="text"
          placeholder="Your last name"
          autocomplete="off"
          required
          v-model.trim="input"
          @keyup="validateInput"
          @blur="validateInput"
          @input="$emit('update:modelValue', $event.target.value)"
      />


    </div>
    <div class="quickErrors" v-if="errors.lastname">
      {{ errors.lastname }}
    </div>
  </div>
</template>
<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";
import MyInput from "@/components/UI/MyInput";

export default {
  name: "LastNameField",
  components: {MyInput},
  setup() {
    let input = ref("");
    const { validateNameField, errors } = useFormValidation();
    const validateInput = () => {
      validateNameField("lastname", input.value);
    };
    return { input, errors, validateInput };
  },
};
</script>
