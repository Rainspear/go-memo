import { Component, Input, OnInit, Output, EventEmitter, OnDestroy } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { faBrain, faClock, faEdit, faPlusCircle, faRotateRight } from '@fortawesome/free-solid-svg-icons';
import { Subscription } from 'rxjs';
import { Memo } from 'src/app/models/memo.model';
import { Schedule } from 'src/app/models/schedule.model';
import { IFilterTopic, Topic } from 'src/app/models/topic.model';
import { ApiService } from 'src/app/services/api.service';
import { MemoDetailService } from 'src/app/services/memo-detail.service';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';
faClock
@Component({
  selector: 'app-topic-detail',
  templateUrl: './topic-detail.component.html',
  styleUrls: ['./topic-detail.component.scss']
})

export class TopicDetailComponent implements OnInit, OnDestroy {
  memos?: Memo[];
  id?: string;
  topic?: Topic;
  schedules?: Schedule[] = []
  topicSubscription?: Subscription;
  paramsSubscription?: Subscription;
  submitSubscription?: Subscription;
  showCreatingMemo: boolean = false;
  showDetailMemo: boolean = false;
  error?: string;

  faClock = faClock;
  faBrain = faBrain;
  faPlusCircle = faPlusCircle;
  faRotateRight = faRotateRight;
  faEdit = faEdit;
  
  constructor(
    private topicSelectingService: TopicSelectingService,
    private apiService: ApiService,
    private route: ActivatedRoute,
    private memoDetailService: MemoDetailService
  ) {
  }

  onToggleCreatingMemo(show: boolean): void {
    this.showCreatingMemo = show;
  }

  onToggleDetailMemo(show: boolean): void {
    this.showDetailMemo = show;
  }

  onSubmitCreatingMemo() {
    this.submitSubscription = this.invokeAllMemo(this?.id || undefined);
  }

  onReloadMemo() {
    this.submitSubscription = this.invokeAllMemo(this?.id || undefined);
  }

  invokeSingleTopic(id: string): Subscription {
    return this.apiService.getSingleTopic(id).subscribe((res: any) => {
      if (res.data) {
        this.topic = res.data;
        this.apiService.getScheduleByTopicId(res.data.id).subscribe((response) => {
          this.schedules = response.data;
        }, (error) => {
          this.error = error.error;
        })
      }
    })
  }

  invokeAllMemo(topic_id?: string): Subscription {
    return this.apiService.getAllMemos(topic_id ? { topic_id } : undefined).subscribe((res: any) => {
      if (res.data) {
        this.memos = res.data;
      }
    }, (error) => {
      this.error = error.error;
    })
  }

  ngOnInit(): void {
    this.id = this.route.snapshot.params['id'];
    if (this.id) this.topicSubscription = this.invokeSingleTopic(this.id);
    this.submitSubscription = this.invokeAllMemo(this?.id || undefined);
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
    if (this.submitSubscription) {
      this.submitSubscription.unsubscribe();
    }
    // if (this.paramsSubscription) {
    //   this.paramsSubscription.unsubscribe();
    // }
  }

}
