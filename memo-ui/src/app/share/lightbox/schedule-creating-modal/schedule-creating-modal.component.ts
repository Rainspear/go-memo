import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { DropdownProps } from 'src/app/constant/dropdown';
import { ParamsCreateSchedule } from 'src/app/models/schedule.model';
import { ApiService } from 'src/app/services/api.service';


@Component({
  selector: 'app-schedule-creating-modal',
  templateUrl: './schedule-creating-modal.component.html',
  styleUrls: ['./schedule-creating-modal.component.scss']
})
export class ScheduleCreatingModalComponent implements OnInit {

  @Input() show: boolean = false;
  @Output() toggle = new EventEmitter<boolean>();
  date: Date = new Date();

  schedule: ParamsCreateSchedule = {
    time: 0,
    level: "essential",
    status: "untouch",
    topic_id: "",
  }

  levelDropdownOptions: DropdownProps[] = [
    {
      title: "Essential",
      value: "essential"
    },
    {
      title: "Critical",
      value: "critical"
    },
    {
      title: "Important",
      value: "important"
    },
    {
      title: "Major",
      value: "major"
    },
    {
      title: "Minor",
      value: "minor"
    },

  ]

  statusDropdownOptions: DropdownProps[] = [
    {
      title: "Untouch",
      value: "untouch"
    },
    {
      title: "Success",
      value: "success"
    },
    {
      title: "Skipped",
      value: "skipped"
    },
    {
      title: "Failure",
      value: "failure"
    },
  ]

  error?: string;

  constructor(private route: ActivatedRoute, private apiService: ApiService) { }

  onToggleHandler(show: boolean) {
    this.toggle.emit(show);
  }

  onSelectLevel(level: DropdownProps) {
    this.schedule.level = level.value;
  }

  onSelectStatus(status: DropdownProps) {
    this.schedule.status = status.value;
  }

  onSubmitHandler() {
    this.schedule.time = Math.floor(new Date(this.date).getTime() / 1000);
    console.log("file: schedule-creating-modal.component.ts ~ line 87 ~ ScheduleCreatingModalComponent ~ onSubmitHandler ~ this.schedule", this.schedule)
    this.apiService.createSchedule(this.schedule).subscribe((res: any) => {
      console.log("ScheduleCreatingModalComponent ~ this.apiService.createSchedule ~ res", res)
      this.onToggleHandler(false)
    }, () => {
      this.onToggleHandler(false)
    })
  }

  ngOnInit(): void {
    this.schedule.topic_id = this.route.snapshot.params['id'];
  }

}
