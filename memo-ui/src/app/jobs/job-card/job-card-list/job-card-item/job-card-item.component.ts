import { Component, Input, OnInit } from '@angular/core';
import { Repetition } from 'src/app/models/topic.model';

@Component({
  selector: 'app-job-card-item',
  templateUrl: './job-card-item.component.html',
  styleUrls: ['./job-card-item.component.scss']
})
export class JobCardItemComponent implements OnInit {

  @Input() job?: Repetition
  constructor() { }

  ngOnInit(): void {
  }

}
