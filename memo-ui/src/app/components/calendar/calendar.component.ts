import { Component, OnInit, Input, OnChanges, SimpleChanges } from '@angular/core';
import { faSquareCaretRight, faSquareCaretLeft } from '@fortawesome/free-regular-svg-icons';
import { faCaretLeft, faCaretRight, faRotateLeft } from '@fortawesome/free-solid-svg-icons';
import { getDaysInMonth, startOfMonth, endOfMonth, isEqual } from 'date-fns';
import { Schedule } from 'src/app/models/schedule.model';

interface DayInWeek {
  name: string;
  value: number;
}

enum FromMonth {
  LastMonth = -1,
  CurrentMonth = 0,
  NextMonth = 1
}

interface DayInView {
  fromMonth: FromMonth;
  value: number;
  date: Date;
  schedules?: Schedule[];
}


@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.scss']
})
export class CalendarComponent implements OnInit, OnChanges {
  faSquareCaretLeft = faSquareCaretLeft;
  faSquareCaretRight = faSquareCaretRight;
  faRotateLeft = faRotateLeft;
  faCaretRight = faCaretRight;
  faCaretLeft = faCaretLeft;

  selectedMonth = new Date().getMonth()
  selectedYear = new Date().getFullYear();
  selectedDay = new Date().getDate();
  selectedDate = this.generateDate();
  daysInView: DayInView[] = this.generateDaysInView();
  daysInWeek: DayInWeek[] = [{
    name: "Sun",
    value: 0
  },
  {
    name: "Mon",
    value: 1
  },
  {
    name: "Tue",
    value: 2
  },
  {
    name: "Wed",
    value: 3
  },
  {
    name: "Thu",
    value: 4
  },
  {
    name: "Fri",
    value: 5
  },
  {
    name: "Sat",
    value: 6
  },
  ];
  scheduleDateInMonth?: Date[];
  @Input() schedules?: Schedule[] = [];

  // firstDayOfMonth = startOfMonth(new Date());
  constructor() {
  }

  convertScheduleToDate(schedules: Schedule[]): Schedule[] {
    return schedules.map(s => {
      s.time_date = new Date(s.time * 1000)
      return s
    })
  }

  mergeSchedulesToDays(schedules: Schedule[]) {
    let scheduleDates = this.convertScheduleToDate(schedules);
    this.daysInView = this.daysInView.map(day => {
      day.schedules = scheduleDates.filter(d => {
        if (d.time_date) {
          return isEqual(
            this.generateDate(day.date.getFullYear(), day.date.getMonth(), day.date.getDate()),
            this.generateDate(d.time_date.getFullYear(), d.time_date.getMonth(), d.time_date.getDate())
          )
        }
        return false;
      }).slice(0, 5) // maximum 5 schedule per day 
      return day;
    })
  }

  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges): void {
    let newSchedules = changes['schedules'].currentValue
    if (newSchedules) {
      this.mergeSchedulesToDays(newSchedules)
    }
  }

  backToCurrentDate() {
    this.selectedMonth = new Date().getMonth()
    this.selectedYear = new Date().getFullYear();
    this.selectedDay = new Date().getDate();
    this.refreshNewValue();
  }

  generateDate(year = this.selectedYear, month = this.selectedMonth, day = this.selectedDay): Date {
    return new Date(`${year}-${this.normalizeMonth(month) + 1}-${day}`);
  }

  normalizeMonth(month: number): number {
    return month > 11 ? 0 : month < 0 ? 11 : month;
  }

  normalizeYear(month: number, year: number = this.selectedYear): number {
    return month > 11 ? year + 1 : month < 0 ? year - 1 : year;
  }

  refreshNewValue() {
    this.selectedDate = this.generateDate();
    this.daysInView = this.generateDaysInView();
    if (this.schedules) this.mergeSchedulesToDays(this.schedules);
  }

  onNextMonthHandler() {
    let month = this.selectedMonth + 1;
    this.selectedMonth = this.normalizeMonth(month);
    this.selectedYear = this.normalizeYear(month);
    this.refreshNewValue();
  }

  onPreviousMonthHandler() {
    let month = this.selectedMonth - 1;
    this.selectedMonth = this.normalizeMonth(month);
    this.selectedYear = this.normalizeYear(month);
    this.refreshNewValue();
  }

  onClickDay(day: DayInView) {
    let month = this.selectedMonth;
    if (day.fromMonth === FromMonth.LastMonth) { month -= 1; }
    if (day.fromMonth === FromMonth.NextMonth) { month += 1; }
    this.selectedMonth = this.normalizeMonth(month);
    this.selectedYear = this.normalizeYear(month)
    this.selectedDay = day.value;
    this.refreshNewValue()
  }

  generateDaysInMonth(year: number = this.selectedYear, month: number = this.selectedMonth, fromMonth: FromMonth = FromMonth.CurrentMonth) {
    let days: DayInView[] = [];
    let today = new Date();
    for (let i = 1; i <= getDaysInMonth(this.generateDate()); i++) {
      days.push({ value: i, fromMonth, date: this.generateDate(year, month, i) })
    }
    return days
  }

  /**
   * Generate DayInView for last, current and next month.
   * @return {DayInView[]} list of DayInView
   */
  generateDaysInView(): DayInView[] {
    // generate days in current month
    let days = this.generateDaysInMonth();
    // generate days in last month
    let firstDayOfMonth = startOfMonth(this.generateDate()).getDay();
    let lastMonth = this.selectedMonth - 1;
    let lastDaysOfLastMonth = endOfMonth(this.generateDate(this.selectedYear, this.normalizeMonth(lastMonth))).getDate();
    for (let i = lastDaysOfLastMonth; i > lastDaysOfLastMonth - firstDayOfMonth; i--) {
      days.unshift(
        {
          value: i,
          fromMonth: FromMonth.LastMonth,
          date: this.generateDate(this.normalizeYear(lastMonth), this.normalizeMonth(lastMonth), i)
        })
    }
    // generate days in next month
    let nextMonth = this.selectedMonth + 1;
    let lastDayOfMonth = endOfMonth(this.generateDate()).getDay();
    let numDaysInNextMonth = days.length < 36 ? 14 : 7; // 
    for (let i = 1; i < numDaysInNextMonth - lastDayOfMonth; i++) {
      days.push({
        value: i,
        fromMonth: FromMonth.NextMonth,
        date: this.generateDate(this.normalizeYear(nextMonth), this.normalizeMonth(nextMonth), i)
      })
    }
    console.log("days", days)
    return days;
  }
}
