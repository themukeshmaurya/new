Migration code :	

 		$table->string('name');
            $table->string('city');
            $table->string('maarks');


===================================

mycontroller.php

<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\facades\DB;

class mycontroller extends Controller
{
    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        $practice = DB::table('practice')->get();
        return view('welcome',['practice'=>$practice]);
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create(Request $request)
    {
        DB::table('practice')->insert([
            'name' =>$request->name,
            'city'=>$request->city,
            'marks'=>$request->marks,
        ]);
        return \redirect(route('index'));
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(Request $request)
    {
        //
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        //
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        $practices=DB::table('practice')->find($id);
        return view('editform',['practices'=>$practices]);
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
        DB::table('practice')->where('id',$id)->update([
            'name' =>$request->name,
            'city'=>$request->city,
            'marks'=>$request->marks,
        ]);
        return \redirect(route('index'));
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        DB::table('practice')->where('id',$id)->delete();
        return \redirect(route('index'));
    }
}

=============================================================

welcome.blade.php


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- CSS only -->
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-rbsA2VBKQhggwzxH7pPCaAqO46MgnOM80zW1RWuH61DGLwZJEdK2Kadq2F9CUG65" crossorigin="anonymous">
<!-- JavaScript Bundle with Popper -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+jjXkk+Q2h455rYXK/7HAuoJl+0I4" crossorigin="anonymous"></script>
    <title>Document</title>
</head>
<body>
    <div class="container mt-5">
        <div class="row">
            <div class="col-sm-6">
            <form action="" method="POST">
                @csrf
                <br/>
                <div class="md-3">
                    <label for="name" class="form-lable">NAME</label>
                    <input type="text" class="form-control" name="name" id="name">
                </div><br/>
                <div class="md-3">
                <label for="city" class="form-lable">CITY</label>
                    <input type="text" class="form-control" name="city" id="city">
                </div><br/>
                <div class="md-3">
                    <label for="marks" class="form-lable">MARKS</label>
                    <input type="text" class="form-control" name="marks" id="marks"></div><br/>
               <button type="submit">SUBMIT</button>
            </form>
            </div class="col-sm-6">

            <div class="col-sm-6">
                <table class="table table-hover">
                    <thead>
                        <tr>
                            <th scope="col">ID</th>
                            <th scope="col">NAME</th>
                            <th scope="col">CITY</th>
                            <th scope="col">MARKS</th>
                            <th scope="col">ACTIONS</th>
                        </tr>
                    </thead>
                    <tbody>
                        @foreach ($practice as $pt)
                        <tr>
                            <th>{{$pt->id}}</th>
                            <td>{{$pt->name}}</td>
                            <td>{{$pt->city}}</td>
                            <td>{{$pt->marks}}</td>
                            <td>
                                <a href="{{url('/edit',$pt->id)}}" class="btn  btn-sm">Edit</a>
                                <a href="{{url('/delete',$pt->id)}}" class="btn  btn-sm">Delete</a>
                            </td>
                        </tr>
                        @endforeach
                    </tbody>
                </table>
            </div class="col-sm-6">
        </div>
    </div>
</body>
</html>


=======================================================

(router)web.php

<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\mycontroller;


Route::get('/',[mycontroller::class,'index'])->name('index');
Route::post('/',[mycontroller::class,'create']);

Route::get('/edit/{id}',[mycontroller::class,'edit'])->name('edit');
Route::post('/edit/{id}',[mycontroller::class,'update'])->name('update');

Route::get('/delete/{id}',[mycontroller::class,'destroy'])->name('destroy');

============================================================================

editform.php


<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- CSS only -->
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-rbsA2VBKQhggwzxH7pPCaAqO46MgnOM80zW1RWuH61DGLwZJEdK2Kadq2F9CUG65" crossorigin="anonymous">
<!-- JavaScript Bundle with Popper -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+jjXkk+Q2h455rYXK/7HAuoJl+0I4" crossorigin="anonymous"></script>
    <title>Document</title>
</head>
<body>
    <div class="container mt-5">
        <div class="row">
            <div class="col-sm-6">
            <form action="" method="POST">
                @csrf
                <br/>
                <div class="md-3">
                    <label for="name" class="form-lable">NAME</label>
                    <input type="text" class="form-control" name="name" id="name" value="{{$practices->name}}">
                </div><br/>
                <div class="md-3">
                <label for="city" class="form-lable">CITY</label>
                    <input type="text" class="form-control" name="city" id="city" value="{{$practices->city}}">
                </div><br/>
                <div class="md-3">
                    <label for="marks" class="form-lable">MARKS</label>
                    <input type="text" class="form-control" name="marks"  id="marks" value="{{$practices->marks}}"></div><br/>
               <button type="submit">Update</button>
            </form>
            </div class="col-sm-6">
        </div>
    </div>
</body>
</html>
