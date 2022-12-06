<template>
  <button class="btn"><slot></slot></button>
</template>

<script>
export default {
  name: "MyButton",
};
</script>

<style scoped>
.btn {
  align-self: flex-end;
  float: right;
  margin: 15px 10px;
  padding: 15px 15px;
  background-color: wheat;
  color: teal;
  border: 1px solid teal;
}
.btn:hover{
  background-color: antiquewhite;
  color: cadetblue;
  box-shadow: 1px 1px 4px gray;
}

.btn:disabled{
  color:gray;
  border: 1px solid gray;
}
</style>
