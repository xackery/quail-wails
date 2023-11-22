

<script setup lang="ts">
//import FollowSpotSetup from "../components/followspot/FollowSpotSetup.vue"
import * as THREE from 'three';
import { start } from 'repl';

const canvas_width = 400;
const canvas_height = 400;

const scene = new THREE.Scene();
function startThree() {
  console.warn('startThree');
  const container = document.getElementById('canvas');
  if (!container) throw new Error('Could not find canvas element');
  document.body.appendChild(container);

  const renderer = new THREE.WebGLRenderer();
  renderer.setPixelRatio(window.devicePixelRatio);
  renderer.setSize(canvas_width, canvas_height);
  container.appendChild(renderer.domElement);

  const camera = new THREE.PerspectiveCamera(45, canvas_width / canvas_height, 1, 1000);
  camera.position.y = 150;
  camera.position.z = 500;
  camera.lookAt(scene.position);
  scene.add(camera);

  scene.add(new THREE.AmbientLight(0x404040));

  //const light = new THREE.PointLight(0xffffff, 1);
  //camera.add(light);

  const geometry = new THREE.BoxGeometry(1, 1, 1);
  const material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
  const cube = new THREE.Mesh(geometry, material);
  scene.add(cube);

  camera.position.z = 5;
  function animate() {
    requestAnimationFrame(animate);
    cube.rotation.x += 0.01;
    cube.rotation.y += 0.01;
    renderer.render(scene, camera);
  }
  animate();
}
</script>

<template>
  <sl-button @Click="startThree">Render</sl-button>

<div id="canvas"></div>
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

