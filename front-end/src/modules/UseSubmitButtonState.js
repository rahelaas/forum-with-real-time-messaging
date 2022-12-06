import { computed } from "vue";

export default function useSubmitButtonState(usere, errors) {

    const isSignupButtonDisabled = computed(() => {
        let disabled = true;
        for (let prop in usere) {
            if (!usere[prop] || errors[prop]) {
                disabled = true;
                break;
            }
            disabled = false;
        }
        return disabled;
    });

    return { isSignupButtonDisabled }
}