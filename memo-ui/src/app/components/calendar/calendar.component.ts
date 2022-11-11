import { Component, OnInit } from '@angular/core';
import { getDaysInMonth, startOfMonth } from 'date-fns'

interface DayInWeek {
  name: string;
  value: number;
}


@Component({
  selector: 'app-calendar',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.scss']
})
export class CalendarComponent implements OnInit {

  daysInMonth: number[] = this.generateDaysInCurrentMonth();
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
  ]
  firstDayOfMonth = startOfMonth(new Date());
  constructor() {
    this.generateDaysInCurrentMonth()
  }

  generateDaysInCurrentMonth() {
    let days : number[] = []
    for (let i = 1; i <= getDaysInMonth(new Date()); i++) {
      days.push(i)
    }
    return days
  }

  ngOnInit(): void {
  }

}
