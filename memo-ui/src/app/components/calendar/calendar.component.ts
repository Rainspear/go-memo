import { Component, OnInit } from '@angular/core';
import {  } from '@fortawesome/free-regular-svg-icons';
import { faRotateLeft, faSquareCaretRight, faSquareCaretLeft } from '@fortawesome/free-solid-svg-icons';
import { getDaysInMonth, startOfMonth, endOfMonth } from 'date-fns';

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
  fromMonth: FromMonth
  value: number;
}


@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.scss']
})
export class CalendarComponent implements OnInit {
  selectedMonth = new Date().getMonth()
  selectedYear = new Date().getFullYear();
  selectedDay = new Date().getDate();
  selectedDate = new Date(this.generateDate());
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

  faSquareCaretLeft = faSquareCaretLeft;
  faSquareCaretRight = faSquareCaretRight;
  faRotateLeft = faRotateLeft;
  // firstDayOfMonth = startOfMonth(new Date());
  constructor() {
  }

  backToCurrentDate() {
    this.selectedMonth = new Date().getMonth()
    this.selectedYear = new Date().getFullYear();
    this.selectedDay = new Date().getDate();
    this.refreshNewValue();
  }

  generateDate(year = this.selectedYear, month = this.selectedMonth, day = this.selectedDay): string {
    return `${year}-${this.normalizeMonth(month) + 1}-${day}`;
  }

  normalizeMonth(month: number) {
    return month > 11 ? 0 : month < 0 ? 11 : month;
  }

  refreshNewValue() {
    this.selectedDate = new Date(this.generateDate());
    this.daysInView = this.generateDaysInView();
  }

  onNextMonthHandler() {
    this.selectedMonth = this.normalizeMonth(this.selectedMonth + 1);
    if (this.selectedMonth === 0) {
      this.selectedYear += 1;
    }
    this.refreshNewValue();
  }

  onPreviousMonthHandler() {
    this.selectedMonth = this.normalizeMonth(this.selectedMonth - 1);
    if (this.selectedMonth === 11) {
      this.selectedYear -= 1;
    }
    this.refreshNewValue();
  }

  onClickDay(day : DayInView) {
    if (day.fromMonth === FromMonth.LastMonth) {
      this.selectedMonth = this.normalizeMonth(this.selectedMonth - 1)
    }
    if (day.fromMonth === FromMonth.NextMonth) {
      this.selectedMonth = this.normalizeMonth(this.selectedMonth + 1)
    }
    this.selectedDay = day.value;
    this.refreshNewValue()
  }

  generateDaysInMonth(fromMonth: FromMonth = FromMonth.CurrentMonth) {
    let days: DayInView[] = [];
    for (let i = 1; i <= getDaysInMonth(new Date(this.generateDate())); i++) {
      days.push({ value: i, fromMonth })
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
    console.log("days", days)
    // generate days in last month
    let firstDayOfMonth = startOfMonth(new Date(this.generateDate())).getDay();
    // let lastMonth = this.normalizeMonth(new Date(this.generateDate()).getMonth()-1);
    let lastDaysOfLastMonth = endOfMonth(new Date(this.generateDate(this.selectedYear, this.normalizeMonth(this.selectedMonth - 1)))).getDate();
    for (let i = lastDaysOfLastMonth; i > lastDaysOfLastMonth - firstDayOfMonth; i--) {
      days.unshift({ value: i, fromMonth: FromMonth.LastMonth })
    }
    // generate days in next month
    let lastDayOfMonth = endOfMonth(new Date(this.generateDate())).getDay();
    let numDaysInNextMonth = days.length < 36 ? 14 : 7 ; // 
    for (let i = 1; i < numDaysInNextMonth - lastDayOfMonth; i++) {
      days.push({ value: i, fromMonth: FromMonth.NextMonth })
    }
    return days;
  }

  ngOnInit(): void {
  }

}
