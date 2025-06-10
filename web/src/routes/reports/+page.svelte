<script>
    import { base } from "$app/paths";
    import {Calendar, DayGrid} from '@event-calendar/core';
    import BackArrowIcon from "$lib/icons/BackArrow.svelte";

    let options = $state({
        view: 'dayGridMonth',
        events: [
            // your list of events
        ],
        headerToolbar: {
            start: "title",
            center: "",
            end: "prev,next"
        }
    });
</script>
<style>
:global {
    .ec-title {
        font-weight: normal;
    }
    .ec-button:focus {
        background-color: var(--bg-3);
    }
}
</style>
<div class="grid page">
    <div class="content" style="padding-top: 2rem;">
    <div>
        <a href={base} class="button faint">
            <BackArrowIcon></BackArrowIcon>
            Back
        </a>
    </div>
<p class="center">Select a week report to view</p>
<div class="no-child-margin-top no-child-clickable-effect">
<Calendar plugins={[DayGrid]} {options} />
</div>
</div>
</div>

