<script>
    import { base } from "$app/paths";
    import {Calendar, DayGrid} from '@event-calendar/core';
    import BackArrowIcon from "$lib/icons/BackArrow.svelte";
    let { data } = $props();
    let events = [];
    data.weeks.forEach((week) => {
        if (week.firstDate && week.lastDate) {
            const [startYear, startMonth, startDate] = week.startDate.split("-");
            const startDate = new Date(startYear, startMonth - 1, startDate);
            /* "start" is the first monday, "first" is the actual first day,
            if they're closed on monday, then first might be a different day,
            but start will still be monday */
            const [firstYear, firstMonth, firstDate] = week.firstDate.split("-");
            const firstDate = new Date(firstYear, firstMonth - 1, firstDate);
            const [lastYear, lastMonth, lastDate] = week.lastDate.split("-");
            const lastDate = new Date(lastYear, lastMonth - 1, lastDate);
            events.push({
                title: `Week of ${[
                    "Jan", "Feb", "Mar", "Apr",
                    "May", "Jun", "Jul", "Aug",
                    "Sep", "Oct", "Nov", "Dec"
                ][startDate.getMonth()]} ${startDate.getDate()}`,
                start: firstDate,
                end: lastDate
            })
        }
    })

    let options = $state({
        view: 'dayGridMonth',
        events: weeks,
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

