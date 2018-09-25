<template>
    <div>
        <DbInfoList></DbInfoList>
    </div>
</template>

<script>
    import DbInfoList from './components/DbInfoList'
    import {mapMutations} from "vuex";

    export default {
        name: 'DbInfo',
        components: {DbInfoList},
        computed: {
        },
        methods: {
        }
    };
</script>
