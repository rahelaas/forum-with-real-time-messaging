<template>

  <div class="field">
    <div class="ui left icon input big">
      <div class="label">
        <label>
          {{ label}}
        </label>
      </div>

      <select v-model="input"
              required
              @blur="validateInput"
              @click="validateInput"
              @change="validateInput"
      >

        <option value="" disabled hidden >Select...</option>
        <!-- for loops you always need also a unique key -->
        <option v-for="(gender) in genders" :key="gender">{{gender}}</option>
      </select>
    </div>
    <div class="quickErrors" v-if="errors.gender">
      {{ errors.gender }}
    </div>
    <!-- -->
  </div>
</template>

<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";

export default {
  name: "GenderField",
  props: ['label', 'modelValue'],
  data() {
    return {
      genders: ['Male', 'Female', 'Transgender', 'Non-binary/non-conforming', 'Prefer Not to Respond']
    }
  },

  setup(props, {emit}) {

    let input = ref("");
    const { validateGenderField, errors } = useFormValidation();
    const validateInput = () => {
      emit('update:modelValue', input.value)
      validateGenderField("gender", input.value);

    };
    return { input, errors, validateInput };
  },

};
</script>




