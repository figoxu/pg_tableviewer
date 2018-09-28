<template>
    <div style="overflow-y:scroll;">
        <TableViewList></TableViewList>
    </div>
</template>


<script>
    import TableViewList from './components/TableViewList'
    import {mapMutations} from "vuex";

    export default {
        name: 'DbInfo',
        components: {TableViewList},
        computed: {
        },
        methods: {
        }
    };
</script>
