import { HttpEvent, HttpHandler, HttpInterceptor, HttpRequest } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthInterceptorService implements HttpInterceptor {

  constructor() { }

  intercept(request: HttpRequest<any>, next: HttpHandler) : Observable<HttpEvent<any>> {
    const modifiedRequest = request.clone({
      headers: request.headers.append('Authorization', `Bearer ${localStorage.getItem('token')}`)
    });
    return next.handle(modifiedRequest)
  }
}
