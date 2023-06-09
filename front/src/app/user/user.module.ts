import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoginComponent } from './login/login.component';
import { SingupComponent } from './singup/singup.component';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';



@NgModule({
  declarations: [LoginComponent, SingupComponent],
  imports: [
    CommonModule,
    FormsModule, ReactiveFormsModule
  ]
})
export class UserModule { }
