<script>
    import { page } from "$app/state";
    import { base } from "$app/paths";
    import BackIcon from "$lib/icons/BackArrow.svelte";
</script>
<div class="grid page" style="margin-top: 4rem;">
    <div class="content">
        <div class="flex">
            <a class="button faint" href={base}>
                <BackIcon></BackIcon>
                Back
            </a>
        </div>
        <h2>{page.status} {page.error.message}</h2>
    </div>
</div>
