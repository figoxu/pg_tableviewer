export default  {
    namespaced: true,
    state: {
        mdfVisible: false,
        addVisible:false,
        mdfResource: {},
    },
    mutations: {
        SET_MDF_VISIBLE: (state, visible) => {
            state.mdfVisible = visible
        },
        SET_ADD_VISIBLE: (state, visible) => {
            state.addVisible = visible
        },
        SET_MDF_RESOURCE: (state, resource) => {
            state.mdfResource = resource
        }
    }
}