<script setup lang="ts">
//import FollowSpotSetup from "../components/followspot/FollowSpotSetup.vue"
import '@shoelace-style/shoelace/dist/components/tree/tree.js';
import '@shoelace-style/shoelace/dist/components/tree-item/tree-item.js';
</script>

<template>
<sl-tree class="tree-with-icons tree">
<sl-tree-item expanded>
<sl-icon name="folder"></sl-icon>
Documents

<sl-tree-item>
    <sl-icon name="folder"> </sl-icon>
    Photos
    <sl-tree-item>
    <sl-icon name="image"></sl-icon>
    birds.jpg
    </sl-tree-item>
    <sl-tree-item>
    <sl-icon name="image"></sl-icon>
    kitten.jpg
    </sl-tree-item>
    <sl-tree-item>
    <sl-icon name="image"></sl-icon>
    puppy.jpg
    </sl-tree-item>
</sl-tree-item>

<sl-tree-item>
    <sl-icon name="folder"></sl-icon>
    Writing
    <sl-tree-item>
    <sl-icon name="file"></sl-icon>
    draft.txt
    </sl-tree-item>
    <sl-tree-item>
    <sl-icon name="file-pdf"></sl-icon>
    final.pdf
    </sl-tree-item>
    <sl-tree-item>
    <sl-icon name="file-bar-graph"></sl-icon>
    sales.xls
    </sl-tree-item>
</sl-tree-item>
</sl-tree-item>
</sl-tree>
</template>


<style scoped>
.tree {
    flex: 1;
    /* justify-content: center; */
    max-width: 600px;
    min-width: 200px;
    height: 100%;
    background-color: var(--middle-pane);
    display: flex;
    overflow: hidden;
}
</style>

