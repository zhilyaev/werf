package deploy

type HelmChartOptions struct {
  Set []string
  Values []string
  DryRun bool
  Debug bool
  Timeout time.Duration
}

func DeployHelmChart(chartDir string, releaseName string, namespace string, opts HelmChartOptions) {
  /*
   * Парсит шаблоны
   * Удаляет старые job-ы (реализация anno dapp/recreate)
   * Проверяем статус релиза и решаем что будем делать install или upgrade
   * Находит все job-hook-и и запускает thread-watcher. Этот тред просто выводит инфу, но
   *  никогда не прерывает процесс деплоя, даже если замечены ошибки в релизе. Kubedog pkg rollout.
   * Запускает helm install/upgrade в основном потоке. Без live-вывода.
   * После завершения helm обязательно ждем завершения треда watcher если он еще не завершился. Функция rollout.TrackJob дает гарантию завершения работы при возникновении ошибок в ресурсах.
   * Выводим блок инфы от helm.
   * Находим все deployment-ы и запускаем в основном потоке rollout.TrackDepolyment... по очереди для каждого ресурса.
   *  rollout.TrackDeployment гарантирует, что err=nil означает что деплоймент готов. Ф-ия повисает до таймаута / ошибки в ресурсах.
   *
   * Ф-ия реализует логику auto_purge_trigger_file.
   * Учесть logsFromTime.
   */
}
