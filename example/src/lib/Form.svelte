<script lang="ts">
  let files;
</script>

<form
  id="form"
  on:submit={async (e) => {
    e.preventDefault();
    
    const formData = new FormData();
    formData.append("file", files[0]);
    const res = await fetch("http://localhost:3000/sendFile", {
      method: "POST",
      body: formData,
    });
    const data = await res.json();
    console.log(data);
  }}
>
  <input class="input file-input" type="file" name="file" bind:files  />
  <button class="button" type="submit">Submit</button>
</form>
