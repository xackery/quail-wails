

<script setup lang="ts">
//import FollowSpotSetup from "../components/followspot/FollowSpotSetup.vue"
import '@shoelace-style/shoelace/dist/components/button/button.js';

import File from '../components/layout/File.vue'
import Tree from '../components/layout/Tree.vue'
import Edit from '../components/layout/Edit.vue'
import Preview from '../components/layout/Preview.vue'
import { ref } from 'vue'

const awesome = ref(true)
</script>

<template>
<File />
<Tree />
<Edit v-if="awesome"/>
<Preview v-else/>
</template>



<style scoped>
.view {
    width: 100%;
    height: 100%;
    background-color: var(--bg-dark);
    display: flex;
    align-items: center;
    overflow: hidden;
}


#canvas {
    background-color: #000;
    width: 200px;
    height: 200px;
    border: 1px solid black;
    margin: 100px;
    padding: 0px;
    position: static; /* fixed or static */
    top: 100px;
    left: 100px;
}
</style>

