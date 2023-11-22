
<script lang="ts" setup>
import { RouterView } from 'vue-router';
import Header from './components/layout/Header.vue'
import Footer from './components/layout/Footer.vue'
import '@shoelace-style/shoelace/dist/themes/dark.css';
import { ref } from 'vue';

import '@shoelace-style/shoelace/dist/themes/dark.css';
import { setBasePath } from '@shoelace-style/shoelace/dist/utilities/base-path';
setBasePath('/shoelace/');

const qrCode = ref();

const isNavCollapse = ref<Boolean>(true)
const isDrawerCollapse = ref<Boolean>(true)

const toggleCollapse = () => {
  isNavCollapse.value = !isNavCollapse.value
}

const toggleDrawer = () => {
  isDrawerCollapse.value = !isDrawerCollapse.value
}
</script>

<template>
  <div className="app">
    <!-- <Header style="widows: 1"/> -->
    <div className="container noselect">
            <div className="routerview noselect">
              <router-view />
            </div>
    </div>
    <Footer v-if="false"/>
</div>
</template>


<style>
.app {
    width: 100vw;
    height: 100vh;
    display: flex;
    flex-direction: column;
    background-color: var(--bg-darker);
    position: absolute;
    top: 0;
    left: 0;
    color:var(--text-color);
}

.routerview {
    background-color: var(--panel-left);
    display: flex;
    flex: 5;
    min-width: 300px;
    overflow: hidden;
}
</style>

<style scoped>
.container {
    /* flex: 1; */
    display: flex;
    justify-content: space-between; /* Add space between left and right drawers and routerview */
    flex-direction: row;
    overflow: hidden;
    /* margin: 0.2%; */

}
</style>