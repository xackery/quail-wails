<script setup lang="ts">
import '@shoelace-style/shoelace/dist/components/icon/icon.js';
import '@shoelace-style/shoelace/dist/components/input/input.js';
</script>

<template>
  <div className="header noselect">
    <div className="left">
        <sl-icon-button name="arrow-left-short" label="Left" style="font-size: 2.5rem;"></sl-icon-button>
    </div>
    <div className="search noselect" style="widows: 2">
       <sl-input className="search-bar notselect" placeholder="Search..."/>
       <sl-icon name="search"></sl-icon>
    </div>
    <div className="right">
        <div className="right-actions">
        </div>
    </div>
  </div>
</template>

<style>
  .header {
    height: 40px; /* Adjust height as needed */
    display: flex;
    align-items: center; /* Center content vertically */
    justify-content: space-between; /* Add space between header elements */
    padding: 10px; /* Add padding for spacing */
    margin-bottom: 5px;
    /*background-color: #838383;*/
    border-radius: 8px;
    margin: 0.2%;

  }

  .search-bar {
    padding-left: 20px;
    padding-right: 30px;
    margin-left: 1px;
    border: 2px none #838383;
    border-top-left-radius: 8px;
    border-bottom-left-radius: 8px;
    justify-self: center;
    background-color: var(--gray-b);
    height: 28px;

    outline: none;
    width: 325px;
    color: #D3D3D3;
    flex: 1;
}

.search {
    display: flex;
    align-items: center;
    border: 0.5px solid var(--gray-9);
    background-color: var(--gray-c);
    border-radius: 10px;
    height: 30px;
    margin-top: 5px;
}

.search:focus-within {
    widows: 0;
    border: 2px solid var(--primary-accent);
    border-radius: 10px;
    background-color:var(--primary-accent);
}

.right {
    flex: 1;
    width: 300px;
    display: flex;
    height: 44px;
    align-items: center;
    margin-left: auto;
    justify-content: space-between;
}

.right-actions {
    display: flex;
    align-items: center;
    margin-left: 20%;
}

.left {
    flex: 1;
    display: flex;
    height: 44px;
    align-items: center;
    margin-right: auto;
    justify-content: left;
}

.window-actions {
    height: 30px;
    width: 150px;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.action {
    width: 44px;
    height: 44px;
    margin-top: 5px;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 5px;
}

.action:last-child {
    margin-right: -10px;
}

.right .cancel:last-child:hover {
    border-radius: 0;
    background-color: #b22222;
}

.action:hover{
    background-color: var(--gray-8)
}

.header-active {
    border: 2px solid var(--primary-accent-dark);
    background-color: var(--gray-8);
}

</style>
