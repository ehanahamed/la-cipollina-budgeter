<script>
import { fade } from "svelte/transition";
let { children, data } = $props();

{#key data.transPageKey}
<div in:fade={{ duration: 140, delay: 140 }} out:fade={{ duration: 140 }}>
    {@render children()}
</div>
{/key}
