<template>
  <div class="field">
    <div class="ui left icon input big">

      <my-input
          type="text"
          placeholder="Your first name"
          autocomplete="off"
          required
          v-model.trim="input"
          @keyup="validateInput"
          @blur="validateInput"
          @input="$emit('update:modelValue', $event.target.value)"
      />


    </div>
      <div class="quickErrors" v-if="errors.firstname">
      {{ errors.firstname }}
    </div>
  </div>
</template>
<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";
import MyInput from "@/components/UI/MyInput";
export default {
  name: "FirstNameField",
  components: {MyInput},
  setup() {
    let input = ref("");
    const { validateNameField, errors } = useFormValidation();
    const validateInput = () => {
      validateNameField("firstname", input.value);
    };
    return { input, errors, validateInput };
  },
};
</script>
