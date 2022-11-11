import { Component, Input, OnInit, Output, EventEmitter, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { Subscription } from 'rxjs';
import { Schedule } from 'src/app/models/schedule.model';
import { IFilterTopic, Topic } from 'src/app/models/topic.model';
import { ApiService } from 'src/app/services/api.service';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';

@Component({
  selector: 'app-topic-detail',
  templateUrl: './topic-detail.component.html',
  styleUrls: ['./topic-detail.component.scss']
})
export class TopicDetailComponent implements OnInit, OnDestroy {
  id?: string;
  topic?: Topic;
  schedules?: Schedule[] = []
  topicSubscription?: Subscription;
  paramsSubscription?: Subscription;
  // @Input() topic?: Topic;

  // @Output() clickClearSelection = new EventEmitter<void>();

  // onClickClearSelection() {
  //   this.clickClearSelection.emit();
  // }
  constructor(
    private topicSelectingService: TopicSelectingService,
    private apiService: ApiService,
    private route: ActivatedRoute) {
    // this.topicSelectingService.selectedTopic.subscribe(topic => {
    //   this.topic = topic;
    // })

  }

  ngOnInit(): void {
    this.id = this.route.snapshot.params['id'];
    if (this.id) this.topicSubscription = this.apiService.getSingleTopic(this.id).subscribe((res: any) => {
      if (res.data) {
        this.topic = res.data;
        this.apiService.getScheduleByTopicId(res.data.id).subscribe((response) => {
          this.schedules = response.data;
        })
      }
    })
    // this.paramsSubscription = this.route.params.subscribe((params: Params) => {
    //   this.id = params['id'];
    // })
    // if (this.id) {
    //   this.topicSelectingService.selectedFilter.subscribe((filter: IFilterTopic) => {
    //     if (this.id) {
    //       this.topicSubscription = this.apiService.getSingleTopic(this.id, filter.value).subscribe((res: any) => {
    //         if (res.data) {
    //           this.topic = res.data
    //         }
    //       })
    //     }
    //   })
    // }
            
  }

  ngOnDestroy(): void {
    // unsubscribe just for example, in reality, angular will automatically unsubscribe for us
    if (this.topicSubscription) {
      this.topicSubscription.unsubscribe();
    }
    // if (this.paramsSubscription) {
    //   this.paramsSubscription.unsubscribe();
    // }
  }

}
