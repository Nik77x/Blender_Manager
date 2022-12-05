import {defineStore} from "pinia";
import Category from "../helpers/Category";


export const  useCurrentCategory = defineStore('category', {

    state: () => ({
        // Default to stable
        category: Category.Library
    })

})