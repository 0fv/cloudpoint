<script setup>
import * as PCL from "pcl.js";
import PointCloudViewer from "pcl.js/PointCloudViewer";
import { onMounted } from "vue";

onMounted(async ()=>{
  const cloudBuffer = await fetch("./result.pcd").then((res) =>
    res.arrayBuffer()
  );
  await PCL.init({
    url:"./pcl-core.wasm"
  })
  
  let cloud = PCL.loadPCDData(cloudBuffer, PCL.PointXYZ);
  const start = Date.now();
  const sor = new PCL.StatisticalOutlierRemoval();
  sor.setInputCloud(cloud);
  sor.setMeanK(40);
  sor.setStddevMulThresh(1.0);
  let cloudFiltered = sor.filter();

  console.log((Date.now() - start) / 1000);

  document.getElementById("progress").style.display = "none";
  document.getElementById("container").style.display = "block";
  showPointCloud(cloud,cloudFiltered);
  bindEvent();
})



  function showPointCloud(cloud,cloudFiltered) {
  let viewer = new PointCloudViewer(
    document.getElementById("canvas"),
    window.innerWidth,
    window.innerHeight
  );
  // viewer.
  viewer.addPointCloud(cloudFiltered);
  viewer.setPointCloudProperties({ color: "#F00" });
  viewer.setAxesHelper({ visible: true, size: 1 });
  viewer.setCameraParameters({ position: { x: 0, y: 0, z: 1.5 } });
  window.addEventListener("resize", () => {
    viewer.setSize(window.innerWidth, window.innerHeight);
  });
}

function bindEvent() {
  const radioOriginal = document.getElementById("original");
  const radioFiltered = document.getElementById("filtered");

  [radioOriginal, radioFiltered].forEach((el) => {
    el.addEventListener("change", (e) => {
      const mode = e.target.id;
      switch (mode) {
        case "original":
          viewer.addPointCloud(cloud);
          viewer.setPointCloudProperties({ color: "#F00" });
          break;
        case "filtered":
          viewer.addPointCloud(cloudFiltered);
          viewer.setPointCloudProperties({ color: "#0F0" });
          break;
      }
    });
  });
}

</script>

<template>
<div class="info">
        <a href="https://pcljs.org" target="_blank" rel="noopener">pcl.js</a>
        - filters - StatisticalOutlierRemoval
      </div>
      <h1
        id="progress"
        style="position: absolute; z-index: 9; right: 50%; top: 50%;"
      >
        Loading...
      </h1>

      <div id="container">
        <canvas id="canvas"></canvas>
        <div
          style="
            position: absolute;
            z-index: 1;
            top: 10px;
            right: 10px;
            background: rgba(0, 0, 0, 0.5);
            color: #fff;
          "
        >
          <fieldset>
            <legend>Select display mode</legend>
            <!-- <div>
            <input type="radio" id="mix" name="display" value="mix" />
            <label for="mix">混合</label>
          </div> -->
            <div>
              <input
                type="radio"
                id="original"
                name="display"
                value="original"
                checked
              />
              <label for="original">Original</label>
            </div>
            <div>
              <input
                type="radio"
                id="filtered"
                name="display"
                value="filtered"
              />
              <label for="filtered">Filtered</label>
            </div>
          </fieldset>
        </div>
      </div>
</template>

<style scoped>
</style>
