import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { faNoteSticky, faCalendar } from '@fortawesome/free-regular-svg-icons';
import { faBrain, faCaretLeft, faClock, faEdit, faPlusCircle, faRotateRight, IconDefinition } from '@fortawesome/free-solid-svg-icons';
import { sub } from 'date-fns';
import { Subscription } from 'rxjs';
import { Memo } from 'src/app/models/memo.model';
import { ParamsFilterSchedule, Schedule } from 'src/app/models/schedule.model';
import { IFilterTopic, Topic } from 'src/app/models/topic.model';
import { ApiService } from 'src/app/services/api.service';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';


interface ITab {
  name: string;
  icon: IconDefinition;
}


@Component({
  selector: 'app-topic-detail',
  templateUrl: './topic-detail.component.html',
  styleUrls: ['./topic-detail.component.scss']
})
export class TopicDetailComponent implements OnInit, OnDestroy {
  faClock = faClock;
  faBrain = faBrain;
  faPlusCircle = faPlusCircle;
  faRotateRight = faRotateRight;
  faEdit = faEdit;
  faCalendar = faCalendar;
  faNoteSticky = faNoteSticky;
  faCaretLeft = faCaretLeft;

  id?: string;

  memos?: Memo[];
  topic?: Topic;
  schedules?: Schedule[]
  filter?: IFilterTopic

  topicSubscription?: Subscription;
  scheduleSubscription?: Subscription;
  memosSubscription?: Subscription;

  showCreatingMemo: boolean = false;
  showDetailMemo: boolean = false;
  showCreatingSchedule: boolean = false;

  tabs: ITab[] = [{ name: "schedules", icon: this.faClock }, { name: "memos", icon: this.faBrain }]
  currentTab: ITab = this.tabs[0];
  selectedTab: "Memo" | "Schedule" = "Schedule";

  error?: string;
  date: Date = new Date();

  constructor(
    private topicSelectingService: TopicSelectingService,
    private apiService: ApiService,
    private route: ActivatedRoute,

  ) {
    this.topicSelectingService.selectedFilter.subscribe((filter: IFilterTopic) => {
      this.filter = filter;
      this.scheduleSubscription = this.invokeSchedules({ topic_id: this.route.snapshot.params['id'], ...this.filter?.value })
    })
  }

  onSelectCategory(tab: "Memo" | "Schedule") {
    this.selectedTab = tab;
  }

  onChangeCurrentTab(tabName: string) {
    let selectedTab = this.tabs.filter(tab => tab.name === tabName)[0]
    this.currentTab = selectedTab;
  }

  onToggleCreatingMemo(show: boolean): void {
    this.showCreatingMemo = show;
  }

  onToggleDetailMemo(show: boolean): void {
    this.showDetailMemo = show;
  }

  onToggleCreatingSchedule(show: boolean): void {
    this.showCreatingSchedule = show;
  }

  onSubmitCreatingMemo() {
    this.memosSubscription = this.invokeAllMemo(this?.id || undefined);
  }

  onReloadMemo() {
    this.memosSubscription = this.invokeAllMemo(this?.id || undefined);
  }

  invokeSchedules(filter: ParamsFilterSchedule): Subscription {
    return this.apiService.getScheduleByFilter(filter).subscribe((res: any) => {
      this.schedules = res.data;
    })
  }

  invokeSingleTopic(id: string): Subscription {
    return this.apiService.getSingleTopic(id).subscribe((res: any) => {
      if (res.data) {
        this.topic = res.data;
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
    if (this.id) this.memosSubscription = this.invokeAllMemo(this.id);
    if (this.id) this.scheduleSubscription = this.invokeSchedules({ topic_id: this.id, ...this.filter?.value })
  }

  ngOnDestroy(): void {
    // unsubscribe just for example, in reality, angular will automatically unsubscribe for us
    if (this.topicSubscription) {
      this.topicSubscription.unsubscribe();
    }
    if (this.memosSubscription) {
      this.memosSubscription.unsubscribe();
    }
    if (this.scheduleSubscription) {
      this.scheduleSubscription.unsubscribe();
    }
  }

}
