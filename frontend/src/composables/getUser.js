import {
    ref
} from "vue";

const user = ref(null)

const getUser = () => {
    return {
        user
    }
}

export default getUser