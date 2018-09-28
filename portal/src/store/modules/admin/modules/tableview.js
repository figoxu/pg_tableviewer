export default {
    namespaced: true,
    state: {
        mdfTableVisible: false,
        mdfColumnVisible:false,
        mdfResource: {},
    },
    mutations: {
        SET_MDF_TABLE_VISIBLE: (state, visible) => {
            state.mdfTableVisible = visible
        },
        SET_MDF_COLUMN_VISIBLE: (state, visible) => {
            state.mdfColumnVisible = visible
        },
        SET_MDF_RESOURCE: (state, resource) => {
            state.mdfResource = resource
        }
    }
}
