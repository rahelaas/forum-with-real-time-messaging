<template>

  <div class="field">
    <div class="ui left icon input big">
      <div class="label">
        <label>
          {{ label}}
        </label>
      </div>

      <select v-model="input"
              required @blur="validateInput"
              @click="validateInput"
              @change="validateInput"
      >

        <option value="" disabled hidden >Select...</option>
        <option :key="year" v-for="year in years" :value="year">
          {{ year }}
        </option>
      </select>



    </div>
    <div class="quickErrors" v-if="errors.year">
      {{ errors.year }}
    </div>
  </div>
</template>

<script>
import { ref } from "vue";
import useFormValidation from "@/modules/useFormValidation";

export default {
  name: "BirthYearField",
  props: ['label', 'modelValue'],
  computed: {

    years() {
      const year = new Date().getFullYear();
      return Array.from(
          { length: year - 1939 },
          (value, index) => 1940 + index
      );
    },
  },


  setup( props, {emit}) {

    let input = ref("");
    const { validateGenderField, errors } = useFormValidation();
    const validateInput = () => {
      emit('update:modelValue', input.value)
      validateGenderField("year", input.value);
    };
    return { input, errors, validateInput };
  },
};
</script>




