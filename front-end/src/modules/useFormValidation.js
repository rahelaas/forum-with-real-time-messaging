import { reactive } from "@vue/reactivity";
import useValidators from '@/modules/Validators'
const errors = reactive({});




export default function useFormValidation() {

    const { isEmpty, minLength, isEmail, isEqual } = useValidators();

    const validateUsernameField = (fieldName, fieldValue) => {
        errors[fieldName] = !fieldValue ? isEmpty(fieldName, fieldValue) : minLength(fieldName, fieldValue, 4)
    }

    const validateNameField = (fieldName, fieldValue) => {
        errors[fieldName] = !fieldValue ? isEmpty(fieldName, fieldValue) : minLength(fieldName, fieldValue, 2)
    }

    const validateEmailField = (fieldName, fieldValue) => {
        errors[fieldName] = !fieldValue ? isEmpty(fieldName, fieldValue) : isEmail(fieldName, fieldValue)
    }
    function validateGenderField (fieldName, fieldValue) {
        errors[fieldName] = !fieldValue ? isEmpty(fieldName, fieldValue) : ''
    }
    const validatePasswordField = (fieldName, fieldValue) => {
        errors[fieldName] = !fieldValue ? isEmpty(fieldName, fieldValue) : minLength(fieldName, fieldValue, 4)
    }

    const validateConfirmPasswordField = (fieldName, fieldValue) => {
        // condition ? truthy : falsy
        errors[fieldName] = !fieldValue ? isEmpty(fieldName, fieldValue) : isEqual(fieldName, fieldValue, 4)
    }

    return { errors, validateUsernameField, validateNameField, validateEmailField, validateGenderField, validatePasswordField, validateConfirmPasswordField }
}


