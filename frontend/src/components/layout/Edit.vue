<script setup lang="ts">
//import FollowSpotSetup from "../components/followspot/FollowSpotSetup.vue"
  import '@shoelace-style/shoelace/dist/components/select/select.js';
  import '@shoelace-style/shoelace/dist/components/option/option.js';
  import '@shoelace-style/shoelace/dist/components/textarea/textarea.js';
  import '@shoelace-style/shoelace/dist/components/input/input.js';
  import '@shoelace-style/shoelace/dist/components/checkbox/checkbox.js';
  import '@shoelace-style/shoelace/dist/components/button/button.js';
/*   const form = document.querySelector('.input-validation-required');

  if (!form) throw new Error('Could not find form element');
  // Wait for controls to be defined before attaching form listeners
  await Promise.all([
    customElements.whenDefined('sl-button'),
    customElements.whenDefined('sl-checkbox'),
    customElements.whenDefined('sl-input'),
    customElements.whenDefined('sl-option'),
    customElements.whenDefined('sl-select'),
    customElements.whenDefined('sl-textarea')
  ]).then(() => {
    form.addEventListener('submit', event => {
      event.preventDefault();
      alert('All fields are valid!');
    });
  }); */
</script>

<template>
<form className="input-validation-required edit">
  <sl-input name="name" label="Name" required></sl-input>
  <br />
  <sl-select label="Favorite Animal" clearable required>
    <sl-option value="birds">Birds</sl-option>
    <sl-option value="cats">Cats</sl-option>
    <sl-option value="dogs">Dogs</sl-option>
    <sl-option value="other">Other</sl-option>
  </sl-select>
  <br />
  <sl-textarea name="comment" label="Comment" required></sl-textarea>
  <br />
  <sl-checkbox required>Check me before submitting</sl-checkbox>
  <br /><br />
  <sl-button type="submit" variant="primary">Submit</sl-button>
</form>

</template>


<style scoped>
.edit {
    flex: 1;
    min-height: 3000px;
    background-color: var(--right-pane);
}
</style>

