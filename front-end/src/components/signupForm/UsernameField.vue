<template>
  <div class="field">
    <div class="ui left icon input big">

      <my-input
          type="text"
          placeholder="Your username"
          autocomplete="off"
          required
          v-model.trim="input"
          @keyup="validateInput"
          @blur="validateInput"
          @input="$emit('update:modelValue', $event.target.value)"
      />

    </div>
    <div class="quickErrors" v-if="errors.username">
      {{ errors.username }}
    </div>
  </div>
</template>
<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";
import MyInput from "@/components/UI/MyInput";
export default {
  name: "UsernameField",
  components: {MyInput},
  setup() {
    let input = ref("");
    const { validateUsernameField, errors } = useFormValidation();
    const validateInput = () => {
      validateUsernameField("username", input.value);
    };
    return { input, errors, validateInput };
  },
};
</script>

<style>

.quickErrors{
  color:red;
}
</style>