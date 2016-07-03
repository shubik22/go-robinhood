const gulp = require('gulp');
const path = require('path');
const del = require('del');
const browserify = require('browserify');
const source = require('vinyl-source-stream');
const uglify = require('gulp-uglify');
const buffer = require('vinyl-buffer');
const flow = require('gulp-flowtype');

const jsSrcPath = path.join('.', 'js');
const assetsPath = path.join('.', 'static_files');
const cssPath = path.join(assetsPath, 'css');
const jsPath = path.join(assetsPath, 'js');

gulp.task('build-css', () => {
  const milligramSrcPath = path.join('.', 'node_modules', 'milligram', 'dist', 'milligram.min.css');
  const normalizeSrcPath = path.join('.', 'node_modules', 'normalize.css', 'normalize.css');

  del(path.join(cssPath, '*'));

  gulp.src(milligramSrcPath)
      .pipe(gulp.dest(cssPath));

  gulp.src(normalizeSrcPath)
      .pipe(gulp.dest(cssPath));
});

gulp.task('typecheck-js', () => {
  const flowConfig = {
    all: false,
    weak: false,
    killFlow: false,
    beep: true,
    abort: false
  };

  return gulp.src(path.join(jsSrcPath, '**', '*.js'))
    .pipe(flow({
      all: false,
      weak: false,
      killFlow: false,
      beep: true,
      abort: false
    }))
});

gulp.task('compile-js', () => {
  const entryPoint = path.join(jsSrcPath, 'app.js');

  del(path.join(jsPath, '*'));

  return browserify(entryPoint)
    .transform("babelify")
    .bundle()
    .pipe(source('app.bundle.js'))
    .pipe(buffer())
    .pipe(uglify())
    .pipe(gulp.dest(jsPath));
});

gulp.task('build-js', ['typecheck-js', 'compile-js']);
gulp.task('build', ['build-css', 'build-js']);

