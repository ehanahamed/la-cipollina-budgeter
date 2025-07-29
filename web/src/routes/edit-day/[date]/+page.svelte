<script>
    import EditDay from "../../EditDay.svelte";
    let { data } = $props();
</script>
<EditDay data={{ ...data, new: false }}></EditDay>
