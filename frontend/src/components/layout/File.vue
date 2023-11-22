<script setup lang="ts">
//import FollowSpotSetup from "../components/followspot/FollowSpotSetup.vue"
import '@shoelace-style/shoelace/dist/components/tree/tree.js';
import '@shoelace-style/shoelace/dist/components/tree-item/tree-item.js';
</script>

<template>
<sl-tree class="tree-with-icons file">
  <sl-tree-item expanded><sl-icon name="folder"></sl-icon>Documents</sl-tree-item>
  <sl-tree-item expanded><sl-icon name="file"></sl-icon>File</sl-tree-item>
  <sl-tree-item expanded><sl-icon name="image"></sl-icon>Image</sl-tree-item>
</sl-tree>
</template>


<style scoped>
.file {
    border: solid 1px var(--sl-color-neutral-200);

    flex: 1;
    justify-content: flex-start;
    max-width: 300px;
    height: 100%;
    background-color: var(--left-pane);
    display: flex;
    overflow: hidden;
}
</style>

