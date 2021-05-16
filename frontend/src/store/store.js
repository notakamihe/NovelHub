import {createStore} from 'vuex';

import { user } from './userStore';

export const store = createStore({
    modules: {
        user
    }
})